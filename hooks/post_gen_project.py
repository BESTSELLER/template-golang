"""
Does the following:

1. Inits git if used
3. Deletes config utils if not needed
"""
from __future__ import print_function
import os
import shutil
from subprocess import Popen
import requests
import os

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

def remove_file(filename):
    """
    generic remove file from project dir
    """
    fullpath = os.path.join(PROJECT_DIRECTORY, filename)
    if os.path.exists(fullpath):
        os.remove(fullpath)

def init_git():
    """
    Initialises git on the new project folder
    """
    GIT_COMMANDS = [
        ["git", "init"],
        ["git", "add", "."],
        ["git", "commit", "-a", "-m", "Initial Commit."]
    ]

    for command in GIT_COMMANDS:
        git = Popen(command, cwd=PROJECT_DIRECTORY)
        git.wait()


def enable_branch_protection(gh_owner, gh_repo, gh_branch):
    url = "https://api.github.com/repos/" + gh_owner + "/" + gh_repo + "/branches/" + gh_branch + "/protection"

    git_hub_token = os.environ['GITHUB_TOKEN']

    if git_hub_token == "":
        print('environmental variable GITHUB_TOKEN needs to be set')
        exit(1)

    jsonfile = open(os.path.curdir +"/hooks/branch-protection.json", "rb")
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




init_git()

enable_branch_protection("BESTSELLER", PROJECT_DIRECTORY, "master")