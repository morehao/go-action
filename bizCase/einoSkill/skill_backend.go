package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/eino/adk/middlewares/skill"
	"gopkg.in/yaml.v3"
)

const skillFileName = "SKILL.md"

// localSkillBackend is a simple skill.Backend implementation that reads
// SKILL.md files from subdirectories under a base directory.
type localSkillBackend struct {
	baseDir string
}

func newLocalSkillBackend(baseDir string) skill.Backend {
	return &localSkillBackend{baseDir: baseDir}
}

func (b *localSkillBackend) List(_ context.Context) ([]skill.FrontMatter, error) {
	skills, err := b.loadAll()
	if err != nil {
		return nil, err
	}
	matters := make([]skill.FrontMatter, 0, len(skills))
	for _, s := range skills {
		matters = append(matters, s.FrontMatter)
	}
	return matters, nil
}

func (b *localSkillBackend) Get(_ context.Context, name string) (skill.Skill, error) {
	skills, err := b.loadAll()
	if err != nil {
		return skill.Skill{}, err
	}
	for _, s := range skills {
		if s.Name == name {
			return s, nil
		}
	}
	return skill.Skill{}, fmt.Errorf("skill %q not found", name)
}

func (b *localSkillBackend) loadAll() ([]skill.Skill, error) {
	entries, err := os.ReadDir(b.baseDir)
	if err != nil {
		return nil, fmt.Errorf("read skills dir: %w", err)
	}

	var skills []skill.Skill
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		skillFile := filepath.Join(b.baseDir, entry.Name(), skillFileName)
		s, err := loadSkillFromFile(skillFile)
		if err != nil {
			return nil, fmt.Errorf("load skill %q: %w", entry.Name(), err)
		}
		s.BaseDirectory = filepath.Join(b.baseDir, entry.Name())
		skills = append(skills, s)
	}
	return skills, nil
}

func loadSkillFromFile(path string) (skill.Skill, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return skill.Skill{}, fmt.Errorf("read file: %w", err)
	}

	fm, content, err := parseFrontmatter(string(data))
	if err != nil {
		return skill.Skill{}, fmt.Errorf("parse frontmatter: %w", err)
	}

	var matter skill.FrontMatter
	if err = yaml.Unmarshal([]byte(fm), &matter); err != nil {
		return skill.Skill{}, fmt.Errorf("unmarshal frontmatter: %w", err)
	}

	return skill.Skill{
		FrontMatter: matter,
		Content:     strings.TrimSpace(content),
	}, nil
}

// parseFrontmatter extracts the YAML front matter and body from a markdown string.
// The file must start with "---", followed by YAML, then "---".
func parseFrontmatter(data string) (frontmatter, content string, err error) {
	const delim = "---"

	scanner := bufio.NewScanner(strings.NewReader(data))

	// expect first line to be "---"
	if !scanner.Scan() || strings.TrimSpace(scanner.Text()) != delim {
		return "", "", fmt.Errorf("file does not start with frontmatter delimiter")
	}

	var fmLines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == delim {
			// collect remaining as body
			var bodyLines []string
			for scanner.Scan() {
				bodyLines = append(bodyLines, scanner.Text())
			}
			return strings.Join(fmLines, "\n"), strings.Join(bodyLines, "\n"), nil
		}
		fmLines = append(fmLines, line)
	}
	return "", "", fmt.Errorf("frontmatter closing delimiter not found")
}
