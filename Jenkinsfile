#!/usr/bin/env groovy

node('docker'){
    String applicationName="jenkins-test"
    String buildNumber="0.1${env.BUILD_NUMBER}"
    String oPath="/go/src/github.com/wenchangshou/${appliationName}"

    stage('Checkout from Github'){
        checkout scm
    }
    stage("Create binaries"){
        docker.image("golang:1.12.5-stretch").inside("-v ${pwd()}:${goPath}"){
            for (command in binaryBuildCommands){
                 sh "cd ${goPath} && GOOS=darwin GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/darwin/${applicationName}-${buildNumber}.darwin.amd64"
                 // build the Windows x64 binary
                 sh "cd ${goPath} && GOOS=windows GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/windows/${applicationName}-${buildNumber}.windows.amd64.exe"
                 // build the Linux x64 binary
                 sh "cd ${goPath} && GOOS=linux GOARCH=amd64 go build -o binaries/amd64/${buildNumber}/linux/${applicationName}-${buildNumber}.linux.amd64" 
            }
        }
    }
    stage("Archive artifacts"){
        archiveArtifacts artifacts: 'binaries/**',fingerprint: true
    }
}