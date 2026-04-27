import { Marked } from 'marked'
import hljs from 'highlight.js'
import 'highlight.js/styles/github.css'
import 'katex/dist/katex.min.css'
import markedKatex from 'marked-katex-extension'

const marked = new Marked()

// KaTeX math formula support ($...$ inline, $$...$$ display)
marked.use(markedKatex({
  throwOnError: false,
  output: 'html',
  nonStandard: true,
}))

// Admonition support: !!! type [title]\ncontent\n!!!
marked.use({
  extensions: [
    {
      name: 'admonition',
      level: 'block',
      start(src: string) {
        return src.match(/^!!!\s/u)?.index
      },
      tokenizer(src: string) {
        const rule = /^!!!\s*(\w+)\s*(.*?)\s*\n([\s\S]*?)^!!!\s*$/m
        const match = rule.exec(src)
        if (!match) return undefined

        const type = match[1].toLowerCase()
        const title = match[2]?.trim() || type.charAt(0).toUpperCase() + type.slice(1)
        const body = match[3].trim()

        const iconMap: Record<string, string> = {
          tip: '&#x1F4A1;',
          warning: '&#x26A0;&#xFE0F;',
          danger: '&#x1F6A8;',
          info: '&#x2139;&#xFE0F;',
          note: '&#x1F4DD;',
          success: '&#x2705;',
        }

        const colorMap: Record<string, { bg: string; border: string }> = {
          tip:     { bg: '#f0fdf4', border: '#22c55e' },
          warning: { bg: '#fffbeb', border: '#f59e0b' },
          danger:  { bg: '#fef2f2', border: '#ef4444' },
          info:    { bg: '#eff6ff', border: '#3b82f6' },
          note:    { bg: '#f8fafc', border: '#94a3b8' },
          success: { bg: '#f0fdf4', border: '#16a34a' },
        }

        const colors = colorMap[type] || colorMap['note']
        const icon = iconMap[type] || iconMap['note']
        const renderedBody = marked.parse(body) as string

        return {
          type: 'admonition',
          raw: match[0],
          html: `<div class="admonition admonition-${type}" style="background:${colors.bg};border-left:4px solid ${colors.border};padding:12px 16px;border-radius:6px;margin:1em 0;"><div class="admonition-title" style="font-weight:600;margin-bottom:6px;color:${colors.border};">${icon} ${title}</div><div class="admonition-content">${renderedBody}</div></div>`,
        }
      },
      renderer(token: any) {
        const html = token.html;
        // 或者
        // renderer(token: { html: string }) {
        //   const { html } = token;
        // }
        return html
      },
    },
  ],
})

// Code syntax highlighting + mermaid
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
