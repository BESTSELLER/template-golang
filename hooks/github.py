import requests
import os


def enable_branch_protection(gh_owner, gh_repo, gh_branch):
    url = "https://api.github.com/repos/" + gh_owner + "/" + gh_repo + "/branches/" + gh_branch + "/protection"

    git_hub_token = os.environ['GITHUB_TOKEN']

    if git_hub_token == "":
        print('environmental variable GITHUB_TOKEN needs to be set')
        exit(1)

    jsonfile = open("./branch-protection.json", "rb")
    headers = {
        'authorization': 'token ' + git_hub_token,
        'content-type': 'application/json',
        'Accept': 'application/vnd.github.loki-preview+json'
    }

    response = requests.request("PUT", url, data=jsonfile, headers=headers)

    if response.status_code < 200 | response.status_code > 299:
        print("status code: {0}\n{1}".format(response.status_code, response.text))
        print("unabled to enable projection on: {0}/{1}:{2}".format(gh_owner, gh_repo, gh_branch))
        exit(1)
    elif response.status_code == 200:
        print("branch protection has been enabled: {0}/{1}:{2}".format(gh_owner, gh_repo, gh_branch))



