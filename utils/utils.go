package utils

func GetFileExtension(language string) string {
	switch language {
	case "csharp":
		return "cs"
	case "typescript":
		return "tsx"
	case "javascript":
		return "jsx"
	default:
		return ""
	}
}

func GetLanguageExtension(language string) string {
	switch language {
	case "csharp":
		return "cs"
	case "typescript":
		return "ts"
	case "javascript":
		return "js"
	default:
		return ""
	}
}
