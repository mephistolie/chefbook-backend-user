{{/*
Expand the name of the chart.
*/}}
{{- define "chefbook-backend-user-service.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "chefbook-backend-user-service.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "chefbook-backend-user-service.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "chefbook-backend-user-service.labels" -}}
helm.sh/chart: {{ include "chefbook-backend-user-service.chart" . }}
{{ include "chefbook-backend-user-service.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/component: service
app.kubernetes.io/part-of: chefbook-backend
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "chefbook-backend-user-service.selectorLabels" -}}
app.kubernetes.io/name: {{ include "chefbook-backend-user-service.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
environment: {{ include "chefbook-backend-user-service.environment" . }}
{{- end }}

{{/*
Choose Environment
*/}}
{{- define "chefbook-backend-user-service.environment" -}}
{{- if eq (tpl "{{ .Values.config.develop }}" .) "true" }}
{{- print "develop" }}
{{- else }}
{{- print "production" }}
{{- end }}
{{- end }}
