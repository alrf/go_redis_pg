import jenkins.*;
import jenkins.model.*;
import javaposse.jobdsl.dsl.DslScriptLoader;
import javaposse.jobdsl.plugin.JenkinsJobManagement;
import com.cloudbees.plugins.credentials.domains.Domain;
import com.cloudbees.plugins.credentials.CredentialsScope;
import com.cloudbees.plugins.credentials.SystemCredentialsProvider;
import com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl;

def jobDslScript = new File('/usr/share/jenkins/ref/init.groovy.d/jobs.groovy')
def jobManagement = new JenkinsJobManagement(System.out, [:], new File('.'))

new DslScriptLoader(jobManagement).runScript(jobDslScript.text)

tempToken = new UsernamePasswordCredentialsImpl(CredentialsScope.GLOBAL,"temp_token", "temp_token", "alrf", "1d2ef8e8-867c-4b3c-a19d-1af350cdeca6")
credentialsStore = Jenkins.instance.getExtensionList(com.cloudbees.plugins.credentials.SystemCredentialsProvider.class)[0];
credentialsStore.store.addCredentials(Domain.global(), tempToken);
