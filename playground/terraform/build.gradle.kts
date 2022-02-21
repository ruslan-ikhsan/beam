/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * License); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an AS IS BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import com.pswidersk.gradle.terraform.TerraformTask
import java.io.ByteArrayOutputStream

plugins {
    id("com.pswidersk.terraform-plugin") version "1.0.0"
}

terraformPlugin {
    terraformVersion.set("1.0.9")
}

tasks {
    /* init Infrastucture for migrate */
    register<TerraformTask>("terraformInit") {
        // exec args can be passed by commandline, for example
        var project_id = project.property("project_id") as String?
        var environment = project.property("project_environment") as String?
        args(
            "init", "-migrate-state",
            "-backend-config=./environment/$environment/state.tfbackend",
            "-var=project_id=$project_id",
            "-var=environment=$environment",
            if (file("./environment/$environment/terraform.tfvars").exists()) {
                "-var-file=./environment/$environment/terraform.tfvars"
            } else {
                "-no-color"
            }
        )
    }
    /* refresh Infrastucture for remote state */
    register<TerraformTask>("terraformRef") {
        mustRunAfter(":playground:terraform:terraformInit")
        var project_id = project.property("project_id") as String?
        var environment = project.property("project_environment") as String?
        args(
            "refresh",
            "-lock=false",
            "-var=project_id=$project_id",
            "-var=environment=$environment",
            if (file("./environment/$environment/terraform.tfvars").exists()) {
                "-var-file=./environment/$environment/terraform.tfvars"
            } else {
                "-no-color"
            }
        )
    }

    /* deploy all App */
    register<TerraformTask>("terraformApplyApp") {
        var project_id = project.property("project_id") as String?
        var environment = project.property("project_environment") as String?
        var docker_tag = if (project.hasProperty("docker-tag")) {
            project.property("docker-tag") as String
        } else {
            environment
        }
        args(
            "apply",
            "-auto-approve",
            "-lock=false",
            "-target=module.applications",
            "-var=project_id=$project_id",
            "-var=environment=$environment",
            "-var=docker_image_tag=$docker_tag",
            if (file("./environment/$environment/terraform.tfvars").exists()) {
                "-var-file=./environment/$environment/terraform.tfvars"
            } else {
                "-no-color"
            }
        )
    }

    /* deploy  App - Only all services for  backend */
    register<TerraformTask>("terraformApplyAppBack") {
        var project_id = project.property("project_id") as String?
        var environment = project.property("project_environment") as String?
        var docker_tag = if (project.hasProperty("docker-tag")) {
            project.property("docker-tag") as String
        } else {
            environment
        }
        args(
            "apply",
            "-auto-approve",
            "-lock=false",
            "-target=module.applications.module.backend",
            "-var=project_id=$project_id",
            "-var=environment=$environment",
            "-var=docker_image_tag=$docker_tag",
            if (file("./environment/$environment/terraform.tfvars").exists()) {
                "-var-file=./environment/$environment/terraform.tfvars"
            } else {
                "-no-color"
            }
        )
    }

    /* deploy  App - Only services for frontend */
    register<TerraformTask>("terraformApplyAppFront") {
        var project_id = project.property("project_id") as String?
        var environment = project.property("project_environment") as String?
        var docker_tag = if (project.hasProperty("docker-tag")) {
            project.property("docker-tag") as String
        } else {
            environment
        }
        args(
            "apply",
            "-auto-approve",
            "-lock=false",
            "-target=module.applications.module.frontend",
            "-var=project_id=$project_id",
            "-var=environment=$environment",
            "-var=docker_image_tag=$docker_tag",
            if (file("./environment/$environment/terraform.tfvars").exists()) {
                "-var-file=./environment/$environment/terraform.tfvars"
            } else {
                "-no-color"
            }
        )
    }

    /* build only Infrastructurte */
    register<TerraformTask>("terraformApplyInf") {
        var project_id = project.property("project_id") as String?
        var environment = project.property("project_environment") as String?
        args(
            "apply",
            "-auto-approve",
            "-lock=false",
            "-target=module.infrastructure",
            "-var=project_id=$project_id",
            "-var=environment=$environment",
            if (file("./environment/$environment/terraform.tfvars").exists()) {
                "-var-file=./environment/$environment/terraform.tfvars"
            } else {
                "-no-color"
            }
        )
    }

    /* build All */
    register<TerraformTask>("terraformApply") {
        var project_id = project.property("project_id") as String?
        var environment = project.property("project_environment") as String?
        var docker_tag = if (project.hasProperty("docker-tag")) {
            project.property("docker-tag") as String
        } else {
            environment
        }
        args(
            "apply",
            "-auto-approve",
            "-lock=false",
            "-var=project_id=$project_id",
            "-var=environment=$environment",
            "-var=docker_image_tag=$docker_tag",
            if (file("./environment/$environment/terraform.tfvars").exists()) {
                "-var-file=./environment/$environment/terraform.tfvars"
            } else {
                "-no-color"
            }
        )
    }

    register<TerraformTask>("terraformDestroy") {
        var project_id = project.property("project_id") as String?
        var environment = project.property("project_environment") as String?
        args(
            "destroy",
            "-auto-approve",
            "-lock=false",
            "-var=project_id=$project_id",
            "-var=environment=$environment",
            if (file("./environment/$environment/terraform.tfvars").exists()) {
                "-var-file=./environment/$environment/terraform.tfvars"
            } else {
                "-no-color"
            }
        )
    }
}


/* set Docker Registry to params from Inf */
task("setDockerRegistry") {
    //get Docker Registry
    dependsOn(":playground:terraform:terraformInit")
    dependsOn(":playground:terraform:terraformRef")
    try {
        var stdout = ByteArrayOutputStream()
        //set Docker Registry
        exec {
            commandLine = listOf("terraform", "output", "docker-repository-root")
            standardOutput = stdout
        }
        project.rootProject.extra["docker-repository-root"] = stdout.toString().trim().replace("\"", "")
    } catch (e: Exception) {
    }
}

task("readState") {
    dependsOn(":playground:terraform:terraformInit")
    dependsOn(":playground:terraform:terraformRef")
}

task("pushBack") {
    dependsOn(":playground:backend:containers:go:dockerTagsPush")
    dependsOn(":playground:backend:containers:java:dockerTagsPush")
    dependsOn(":playground:backend:containers:python:dockerTagsPush")
    dependsOn(":playground:backend:containers:scio:dockerTagsPush")
    dependsOn(":playground:backend:containers:router:dockerTagsPush")
}

task("pushFront") {
    dependsOn(":playground:frontend:dockerTagsPush")
}

task("prepareConfig") {
    doLast {
        var playgroundBackendUrl = ""
        var playgroundBackendJavaRouteUrl = ""
        var playgroundBackendGoRouteUrl = ""
        var playgroundBackendPythonRouteUrl = ""
        var playgroundBackendScioRouteUrl = ""
        var stdout = ByteArrayOutputStream()
        exec {
            commandLine = listOf("terraform", "output", "router-server-url")
            standardOutput = stdout
        }
        playgroundBackendUrl = stdout.toString().trim().replace("\"", "")
        stdout = ByteArrayOutputStream()
        exec {
            commandLine = listOf("terraform", "output", "go-server-url")
            standardOutput = stdout
        }
        playgroundBackendGoRouteUrl = stdout.toString().trim().replace("\"", "")
        stdout = ByteArrayOutputStream()
        exec {
            commandLine = listOf("terraform", "output", "java-server-url")
            standardOutput = stdout
        }
        playgroundBackendJavaRouteUrl = stdout.toString().trim().replace("\"", "")
        stdout = ByteArrayOutputStream()
        exec {
            commandLine = listOf("terraform", "output", "python-server-url")
            standardOutput = stdout
        }
        playgroundBackendPythonRouteUrl = stdout.toString().trim().replace("\"", "")
        stdout = ByteArrayOutputStream()
        exec {
            commandLine = listOf("terraform", "output", "scio-server-url")
            standardOutput = stdout
        }
        playgroundBackendScioRouteUrl = stdout.toString().trim().replace("\"", "")
        val configFileName = "gradle.properties"
        val modulePath = project(":playground:frontend").projectDir.absolutePath
        var file = File(modulePath + "/" + configFileName)

        file.writeText(
            """################################################################################
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################

playgroundBackendUrl=${playgroundBackendUrl}
analyticsUA=UA-73650088-1
playgroundBackendJavaRouteUrl=${playgroundBackendJavaRouteUrl}
playgroundBackendGoRouteUrl=${playgroundBackendGoRouteUrl}
playgroundBackendPythonRouteUrl=${playgroundBackendPythonRouteUrl}
playgroundBackendScioRouteUrl=${playgroundBackendScioRouteUrl}
"""
        )
        try {
            var stdout = ByteArrayOutputStream()
            //set Docker Registry
            exec {
                commandLine = listOf("terraform", "output", "docker-repository-root")
                standardOutput = stdout
            }
            project.rootProject.extra["docker-repository-root"] = stdout.toString().trim().replace("\"", "")
        } catch (e: Exception) {
        }
    }
}
/* initialization infrastructure */
task("InitInfrastructure") {
    val init = tasks.getByName("terraformInit")
    val apply = tasks.getByName("terraformApplyInf")
    dependsOn(init)
    dependsOn(apply)
    apply.mustRunAfter(init)
}

/* build, push, deploy Frontend app */
task("deployFrontend") {
    val read = tasks.getByName("readState")
    val prepare = tasks.getByName("prepareConfig")
    val push = tasks.getByName("pushFront")
    val deploy = tasks.getByName("terraformApplyAppFront")
    dependsOn(read)
    Thread.sleep(10)

    prepare.mustRunAfter(read)
    dependsOn(prepare)
    push.mustRunAfter(prepare)
    deploy.mustRunAfter(push)
    dependsOn(push)
    dependsOn(deploy)
}

/* build, push, deploy Backend app */
task("deployBackend") {

    val config = tasks.getByName("setDockerRegistry")
    val push = tasks.getByName("pushBack")
    val deploy = tasks.getByName("terraformApplyAppBack")
    dependsOn(config)
    Thread.sleep(10)
    push.mustRunAfter(config)
    deploy.mustRunAfter(push)
    dependsOn(push)
    dependsOn(deploy)

}

