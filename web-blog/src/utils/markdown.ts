import { Marked } from 'marked'
import hljs from 'highlight.js'
import 'highlight.js/styles/github.css'

const marked = new Marked()

marked.use({
  renderer: {
    code({ text, lang }: { text: string; lang?: string }) {
      const language = lang && hljs.getLanguage(lang) ? lang : undefined
      const highlighted = language
        ? hljs.highlight(text, { language }).value
        : hljs.highlightAuto(text).value

      if (lang === 'mermaid') {
        return `<pre class="mermaid">${text}</pre>`
      }

      return `<pre><code class="hljs language-${lang || ''}">${highlighted}</code></pre>`
    },
  },
})

export function renderMarkdown(content: string): string {
  if (!content) return ''
  return marked.parse(content) as string
}
