import requests
import json

class Cdn():

    def __init__(self, email, key):
        self.email = email
        self.key = key

    def setup_caching_level(self, option):

        #replace option with aggressive|basic|simplified

        url = "https://api.cloudflare.com/client/v4/zones/8517bd991c17a1808509cd0d9d31c282/settings/cache_level"
        headers = {
            "X-Auth-Email": self.email,
            "X-Auth-Key": self.key,
            "Content-Type": "application/json",
        }

        data = {"value":option}

        response = requests.patch(url, headers=headers, data=json.dumps(data))

        if response.status_code == 200:
             print("cache level changed successfully")
        else:
             print("fail to change cache parameter")

    def setup_browser_cache_ttl(self, time):

        #replace time with 1800|300|900|1800|2700|3600|7200|10800|14400|28800|57600|86400|604800|2592000|31536000

        url = "https://api.cloudflare.com/client/v4/zones/8517bd991c17a1808509cd0d9d31c282/settings/browser_cache_ttl"

        headers = {
            "X-Auth-Email": self.email,
            "X-Auth-Key": self.key,
            "Content-Type": "application/json",
        }

        data = {"value":time}

        response = requests.patch(url, headers=headers, data=json.dumps(data))

        if response.status_code == 200:
            print("TTL changed successfully")
        else:
            print("fail to change TTL parameter")

    def always_online(self, option):

        #replace option with on|off

        url = "https://api.cloudflare.com/client/v4/zones/8517bd991c17a1808509cd0d9d31c282/settings/always_online"

        headers = {
            "X-Auth-Email": self.email,
            "X-Auth-Key": self.key,
            "Content-Type": "application/json",
        }

        data = {"value":option}

        response = requests.patch(url, headers=headers, data=json.dumps(data))

        if response.status_code == 200:
            print("Always online configured")
        else:
            print("Failed to change parameter")

    def purge_all_files(self, option="true"):

        url = "DELETE https://api.cloudflare.com/client/v4/zones/:identifier/purge_cache"

        headers = {
            "X-Auth-Email": self.email,
            "X-Auth-Key": self.key,
            "Content-Type": "application/json",
        }

        data = {"purge_everything": option}

        response = requests.patch(url, headers=headers, data=json.dumps(data))

        if response.status_code == 200:
            print("files purged!")
        else:
            print("Failed to purge files")


