package dag

type pipelineResult struct {
	dag *Dag
}

func (result *pipelineResult) Then() *pipelineDSL {
	return &pipelineDSL{
		result.dag,
	}
}

func (result *pipelineResult) OnComplete(action func() error) *pipelineResult {
	job := result.dag.lastJob()
	if job != nil {
		job.onComplete = action
	}
	return result
}

type pipelineDSL struct {
	dag *Dag
}

func (dsl *pipelineDSL) Spawns(tasks ...func() error) *spawnsResult {
	dsl.dag.Spawns(tasks...)
	return &spawnsResult{
		dsl.dag,
	}
}

type spawnsResult struct {
	dag *Dag
}

func (result *spawnsResult) Join() *spawnsDSL {
	return &spawnsDSL{
		result.dag,
	}
}

func (result *spawnsResult) OnComplete(action func() error) *spawnsResult {
	job := result.dag.lastJob()
	if job != nil {
		job.onComplete = action
	}
	return result
}

type spawnsDSL struct {
	dag *Dag
}

func (dsl *spawnsDSL) Pipeline(tasks ...func() error) *pipelineResult {
	dsl.dag.Pipeline(tasks...)
	return &pipelineResult{
		dsl.dag,
	}
}
