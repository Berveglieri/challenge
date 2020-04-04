import os
import subprocess
import shutil
import git


class Builder():


    def docker_build(self, tag, repo_url, dest_folder, ecr_repo):

        try:
            git.Repo.clone_from(repo_url, dest_folder)
        except Exception as e:
            print(e)

        try:
            os.chdir(dest_folder+"/hive/hive")
            dockerfile = os.path.isfile(os.path.join(os.getcwd(),'Dockerfile'))
            if dockerfile:
                call = subprocess.Popen(['docker', 'build', '-t', ecr_repo+":"+tag, '.'],
                                        stdout=subprocess.PIPE,
                                        stderr=subprocess.PIPE)
                stdout, stderr = call.communicate()
                print(stdout.decode('UTF-8'), stderr.decode('UTF-8'))
                shutil.rmtree(dest_folder)
        except FileNotFoundError as e:
            print(e)
