package sdk_utils

import "strings"

func StripTags(v string) string {
	var buf strings.Builder
	buf.Grow(len(v))

	for i := 0; i < len(v); {
		// Find next '<'
		lt := strings.IndexByte(v[i:], '<')
		if lt == -1 {
			buf.WriteString(v[i:])
			break
		}

		buf.WriteString(v[i : i+lt])
		i += lt // i now points to '<'

		// Need at least one more character
		if i+1 >= len(v) {
			buf.WriteByte('<')
			break
		}

		// Tags must start with <letter, </, <!, or <?
		next := v[i+1]
		if !isAlpha(next) && next != '/' && next != '!' && next != '?' {
			buf.WriteByte('<')
			i++
			continue
		}

		// Find the closing '>', respecting quotes
		gt := findCloseBracket(v, i+1)
		if gt == -1 {
			buf.WriteByte('<')
			i++
			continue
		}

		tagContent := v[i+1 : gt]
		trimmedContent := strings.TrimLeft(tagContent, " \t\r\n")

		// Handle HTML comments: <!-- ... -->
		if len(trimmedContent) > 0 && trimmedContent[0] == '!' {
			_ = tagContent // tagContent used implicitly via trimmedContent
			if strings.HasPrefix(trimmedContent, "!--") {
				// Find closing -->
				closeComment := strings.Index(v[gt+1:], "-->")
				if closeComment != -1 {
					i = gt + 1 + closeComment + 3
				} else {
					i = gt + 1
				}
			} else {
				// Other declarations like <!DOCTYPE>
				i = gt + 1
			}
			continue
		}

		// Handle PHP / processing instructions: <? ... ?>
		// gt already points to the '>' in '?>', so skip past it
		if len(trimmedContent) > 0 && trimmedContent[0] == '?' {
			i = gt + 1
			continue
		}

		// Extract tag name
		tagName := extractTagNameForStrip(tagContent)
		i = gt + 1

		// Script and style — also strip their content
		lowerName := strings.ToLower(tagName)
		if lowerName == "script" || lowerName == "style" {
			closeTag := "</" + lowerName + ">"
			closePos := strings.Index(strings.ToLower(v[i:]), closeTag)
			if closePos != -1 {
				i += closePos + len(closeTag)
			} else {
				i = len(v)
			}
		}
		// For other tags, just skip the tag itself and keep content
	}

	return buf.String()
}

// isAlpha checks if a byte is an ASCII letter.
func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// findCloseBracket finds the matching '>' for a tag starting at start,
// correctly handling attribute values in quotes.
func findCloseBracket(s string, start int) int {
	inQuote := false
	quoteChar := byte(0)
	for i := start; i < len(s); i++ {
		ch := s[i]
		if inQuote {
			if ch == quoteChar {
				inQuote = false
			}
			continue
		}
		if ch == '"' || ch == '\'' {
			inQuote = true
			quoteChar = ch
			continue
		}
		if ch == '>' {
			return i
		}
	}
	return -1
}

// extractTagNameForStrip extracts the tag name from tag content like "a href=..." or "/a".
func extractTagNameForStrip(tagContent string) string {
	i := 0
	// Skip leading whitespace and possible '/'
	for i < len(tagContent) && (tagContent[i] == ' ' || tagContent[i] == '\t' || tagContent[i] == '\n' || tagContent[i] == '\r') {
		i++
	}
	if i < len(tagContent) && tagContent[i] == '/' {
		i++
	}
	// Extract the tag name
	start := i
	for i < len(tagContent) && (isAlpha(tagContent[i]) || (tagContent[i] >= '0' && tagContent[i] <= '9') || tagContent[i] == '-' || tagContent[i] == ':') {
		i++
	}
	if start == i {
		return ""
	}
	return tagContent[start:i]
}
