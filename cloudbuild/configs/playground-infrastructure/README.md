# 

**This README explains how to:**

1. Enable Cloud Build in your GCP Project

2. Connect Github repository to Cloud Build


### Before you begin

- Create GCP project
    - [Create a project docs](https://cloud.google.com/apis/docs/getting-started#creating_a_google_project)
- Enable billing for the project
    - [Enable billing docs](https://cloud.google.com/apis/docs/getting-started#enabling_billing)
- Enable Cloud Build API
    - [Enable API docs](https://cloud.google.com/apis/docs/getting-started#enabling_apis)
- Have your GitHub repository ready

## Enable Cloud Build and connect a GitHub Repository:

To connect your GitHub repository to your Cloud Build:

1. Open the Cloud Build, either from Navigation Menu on the left side or using search engine.

2. Open the Triggers page in the Cloud Build console. Select required GCP project in the top bar.

3. Click Connect repository on the top. Select GitHub (Cloud Build GitHub App), check the consent checkbox, and click Continue.

4. In the pop-up that appears, select your GitHub username or organization.

5. Select one of the following options based on your business need:

        a. All repositories - enable current and future GitHub repositories for access via the Cloud Build app

        b. Only select repositories - use the Select repositories drop-down to enable only specific repositories for access via the Cloud Build app.

6. Click Install to install the Cloud Build app.

7. After you have selected a project or created a new one, you will see the Connect repository panel.

8. In the Select repository section, select the following fields:

* GitHub account: The GitHub account used to install the Cloud Build GitHub App. This field may be pre-selected for you.

* Repository: The repositories you want to connect to Cloud Build.

9. Once you have selected your GitHub account and repositories, read the consent disclaimer and select the checkbox next to proceed further.

10. Click Connect and Done.

You have now connected one or more GitHub repositories to your Cloud project and Cloud Build. 