# -*- coding:utf-8 -*-
"""
iptv-org数据采集程序

https://github.com/iptv-org/api
"""
import mariadb
import requests


class IPTVSpider:

    def __init__(self):
        self.apis = [
            {
                "table": "iptv_channels",
                "api": "https://iptv-org.github.io/api/channels.json",
            },
            "https://iptv-org.github.io/api/streams.json",
            "https://iptv-org.github.io/api/guides.json",
            "https://iptv-org.github.io/api/categories.json",
            "https://iptv-org.github.io/api/languages.json",
            "https://iptv-org.github.io/api/countries.json",
            "https://iptv-org.github.io/api/subdivisions.json",
            "https://iptv-org.github.io/api/regions.json",
        ]
        self.conn = mariadb.connect(
            user="root",
            password="123456",
            host="127.0.0.1",
            port=3307,
            database="watchtv"
        )

    def get_api_data(self, api: str):
        """
        获取接口数据
        :param str api: 接口地址
        :return:
        """
        response = requests.get(api)
        return response.json()


if __name__ == "__main__":
    o = IPTVSpider()
