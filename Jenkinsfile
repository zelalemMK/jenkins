pipeline {
    agent { docker { image 'golang:1.17.5-alpine' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh '''
                echo "whats up bitches"
                '''
            }
        }
        stage('wap') {
            steps {
                retry(3) {
                    sh './wap.sh'
                }
                timeout(time:3, unit:'MINUTES') {
                    sh './wap.sh'
                }
            }
        }
    }
    post {
        always {
                echo "always Ill remember you"
            }
        success {
                echo "success"
            }
        }
}