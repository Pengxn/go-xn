package cron

import (
	"context"
	"errors"
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	// c is the cron instance.
	c *cron.Cron
	// cmap is the cron map, used to store the mapping of cron name to cron entry id.
	// The key is the cron name, and the value is the cron entry id.
	cmap sync.Map
)

func init() {
	c = cron.New()
	cmap = sync.Map{}
}

// AddFunc adds a func to the cron to be run.
func AddFunc(ctx context.Context, spec string, cmd func()) error {
	return AddJob(ctx, spec, cron.FuncJob(cmd))
}

// AddJob adds a Job to the cron to be run, the job implements a Run() method.
func AddJob(ctx context.Context, spec string, cmd cron.Job) error {
	name := ctx.Value("cron").(string)
	if name == "" {
		return errors.New("add cron job failed: cron name is empty")
	}

	entryID, err := c.AddJob(spec, cmd)
	if err != nil {
		return err
	}
	cmap.Store(name, entryID)
	return nil
}

// ListJobs lists all the jobs and entry ids in the `cmap`.
func ListJobs() map[string]int {
	jobs := make(map[string]int)

	cmap.Range(func(key, value any) bool {
		jobs[key.(string)] = int(value.(cron.EntryID))
		return true
	})
	return jobs
}

// RemoveJob removes a job from the cron.
func RemoveJob(ctx context.Context) error {
	name := ctx.Value("cron").(string)
	if name == "" {
		return errors.New("remove cron job failed: cron name is empty")
	}

	entryID, ok := cmap.Load(name)
	if !ok {
		return errors.New("remove cron job failed: cron name not found")
	}
	c.Remove(cron.EntryID(entryID.(int)))
	cmap.Delete(name)
	return nil
}
