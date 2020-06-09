{{ define "packages" }}
---
title: PaCRD CRD Documentation
linkTitle: PaCRD CRD Doc
weight: 172
---

{{ "{{< include 'experimental-feature.html' >}}" }}

{{ range .packages }}
# {{ packageDisplayName . -}}
    {{ with (index .GoPackages 0 )}}
        {{ with .DocComments }}
        <p>
            {{ safe (renderComments .) }}
        </p>
        {{ end }}
    {{ end }}
## Resource Types:
    <ul>
    {{- range (visibleTypes (sortedTypes .Types)) -}}
        {{ if isExportedType . -}}
        <li>
            <a href="{{ linkForType . }}">{{ typeDisplayName . }}</a>
        </li>
        {{- end }}
    {{- end -}}
    </ul>

    {{ range (visibleTypes (sortedTypes .Types))}}
        {{ template "type" .  }}
    {{ end }}

{{ end }}

<p><em>
    Generated with <code>gen-crd-api-reference-docs</code>
    {{ with .gitCommit }} on git commit <code>{{ . }}</code>{{end}}.
</em></p>

{{ end }}
