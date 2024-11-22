package autograder

type Output interface {
	OutputFile()
}

type ImageOutput struct {
}

type TextOutput struct {
}

type FileCheckOutput struct {
}

type CompilationOutput struct {
}
