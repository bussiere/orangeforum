// Copyright (c) 2017 Sagar Gubbi. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const profiletopicsSrc = `
{{ define "content" }}

<h2>Topics by {{ .OwnerName }}</h2>

{{ if .Topics }}
{{ range .Topics }}
<div class="row">
	<div>
		<a href="{{ if not .IsDeleted }}/topics?id={{ .ID }}{{ else }}/topics/edit?id={{ .ID }}{{ end }}">{{ .Title }}{{ if .IsClosed }} [closed]{{ end }}{{ if .IsDeleted }} [deleted]{{ end }}</a>
	</div>
	<div class="muted">{{ .CreatedDate }}</div>
</div>
{{ end }}
{{ else }}
<div class="row">
	<div class="muted">No topics to show.</div>
</div>
{{ end }}

{{ if .LastTopicDate }}
<div class="row">
	<div>
		<a href="/users/topics?u={{ .OwnerName }}&ltd={{ .LastTopicDate }}">More</a>
	</div>
</div>
{{ end }}

{{ end }}
`
