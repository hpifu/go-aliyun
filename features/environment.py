#!/usr/bin/env python3

import pymysql
import redis
import subprocess
import time
import requests
import datetime
import json
import socket
from behave import *


register_type(int=int)
register_type(str=lambda x: x if x != "N/A" else "")
register_type(bool=lambda x: True if x == "true" else False)


config = {
    "prefix": "output/go-aliyun",
    "service": {
        "port": 17060,
        "cookieSecure": False,
        "allowOrigins": ["http://127.0.0.1:4000"],
        "cookieDomain": "127.0.0.1"
    },
}


def wait_for_port(port, host="localhost", timeout=5.0):
    start_time = time.perf_counter()
    while True:
        try:
            with socket.create_connection((host, port), timeout=timeout):
                break
        except OSError as ex:
            time.sleep(0.01)
            if time.perf_counter() - start_time >= timeout:
                raise TimeoutError("Waited too long for the port {} on host {} to start accepting connections.".format(
                    port, host
                )) from ex


def deploy():
    fp = open("{}/configs/aliyun.json".format(config["prefix"]))
    cf = json.loads(fp.read())
    fp.close()
    cf["service"]["port"] = ":{}".format(config["service"]["port"])
    cf["service"]["cookieSecure"] = config["service"]["cookieSecure"]
    cf["service"]["cookieDomain"] = config["service"]["cookieDomain"]
    cf["service"]["allowOrigins"] = config["service"]["allowOrigins"]
    print(cf)
    fp = open("{}/configs/aliyun.json".format(config["prefix"]), "w")
    fp.write(json.dumps(cf, indent=4))
    fp.close()


def start():
    subprocess.Popen(
        "cd {} && nohup bin/aliyun &".format(config["prefix"]),  shell=True
    )

    wait_for_port(config["service"]["port"], timeout=5)


def stop():
    subprocess.getstatusoutput(
        "ps aux | grep bin/aliyun | grep -v grep | awk '{print $2}' | xargs kill"
    )


def before_all(context):
    config["url"] = "http://127.0.0.1:{}".format(config["service"]["port"])
    deploy()
    start()
    context.config = config


def after_all(context):
    stop()
