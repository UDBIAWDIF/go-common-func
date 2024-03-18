package utils

type Runner struct {
	Runnable bool
}

type RunnerFunc func() bool
type RunnerFinishFunc func()

func NewRunner() *Runner {
	runner := new(Runner)
	runner.Runnable = true
	return runner
}

func (_this *Runner) IsRunnable() bool {
	return _this.Runnable
}

func (_this *Runner) Exec(runnerFunc RunnerFunc) *Runner {
	if _this.IsRunnable() {
		_this.Runnable = runnerFunc()
	}
	return _this
}

func (_this *Runner) Success(runnerFunc RunnerFinishFunc) *Runner {
	if _this.IsRunnable() {
		runnerFunc()
	}
	return _this
}

func (_this *Runner) Failed(runnerFunc RunnerFinishFunc) *Runner {
	if !_this.IsRunnable() {
		runnerFunc()
	}
	return _this
}
