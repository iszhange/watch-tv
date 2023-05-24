# -*- coding:utf-8 -*-
"""
iptv-org数据采集程序

https://github.com/iptv-org/api
"""
import datetime
import json

import mariadb
import requests


class IPTVSpider:

    def __init__(self):
        self.proxies = {
            "http": "http://127.0.0.1:7890",
            "https": "http://127.0.0.1:7890"
        }
        self.apis = [
            # {"table": "iptv_channels", "api": "https://iptv-org.github.io/api/channels.json"},
            {"table": "iptv_streams", "api": "https://iptv-org.github.io/api/streams.json"},
            # {"table": "iptv_guides", "api": "https://iptv-org.github.io/api/guides.json"},
            # {"table": "iptv_categories", "api": "https://iptv-org.github.io/api/categories.json"},
            # {"table": "iptv_languages", "api": "https://iptv-org.github.io/api/languages.json"},
            # {"table": "iptv_countries", "api": "https://iptv-org.github.io/api/countries.json"},
            # {"table": "iptv_subdivisions", "api": "https://iptv-org.github.io/api/subdivisions.json"},
            # {"table": "iptv_regions", "api": "https://iptv-org.github.io/api/regions.json"},
        ]
        self.tabels = {
            "iptv_channels": ["id", "name", "alt_names", "network", "owners", "country", "subdivision", "city",
                              "broadcast_area", "languages", "categories", "is_nsfw", "launched", "closed",
                              "replaced_by", "website", "logo"],
            "iptv_streams": ["channel", "url", "http_referrer", "user_agent"],
            "iptv_guides": ["channel", "site", "lang", "days", "url"],
            "iptv_categories": ["id", "name"],
            "iptv_languages": ["name", "code"],
            "iptv_countries": ["name", "code", "languages", "flag"],
            "iptv_subdivisions": ["country", "name", "code"],
            "iptv_regions": ["code", "name", "countries"],
        }
        self.json_fields = {
            "iptv_channels": ["alt_names", "owners", "broadcast_area", "languages", "categories"],
            "iptv_streams": [],
            "iptv_guides": [],
            "iptv_categories": [],
            "iptv_languages": [],
            "iptv_countries": ["languages"],
            "iptv_subdivisions": [],
            "iptv_regions": ["countries"],
        }
        self.conn = mariadb.connect(
            user="root",
            password="123456",
            host="127.0.0.1",
            port=3307,
            database="watchtv"
        )

    def run(self):
        """
        执行
        :return:
        """
        for item in self.apis:
            result = self.get_api_data(item["api"])
            print(result)
            self.save_data(item["table"], result)

    def save_data(self, table: str, dataset: list):
        """
        保存数据到数据库
        :param table:
        :param dataset:
        :return:
        """
        cursor = self.conn.cursor()

        # 清除表数据
        cursor.execute("TRUNCATE " + table)
        self.conn.commit()

        # 插入数据
        fields = self.tabels[table]
        fields.append("created_at")
        fields.append("updated_at")
        field_str = ",".join(fields)
        json_fields = self.json_fields[table]
        field_mark = []
        for i in range(len(fields)):
            field_mark.append("?")
        field_mark_str = ",".join(field_mark)
        sql = "INSERT INTO %s (%s) VALUES (%s)" % (table, field_str, field_mark_str)
        values = []
        i = 0
        for item in dataset:
            print(item)
            value = []
            date = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
            for key in fields:
                if key == "updated_at" or key == "created_at":
                    value.append(date)
                    continue
                v = item[key]
                if key == "is_nsfw":
                    value.append(str(int(v)))
                    continue
                if key in json_fields:
                    value.append("[]" if v is None else json.dumps(v, ensure_ascii=False))
                else:
                    value.append("" if v is None else v)
            values.append(value)
            # 每100条入一次库
            if i > 1:
                try:
                    cursor.executemany(sql, values)
                    self.conn.commit()
                    i = 0
                    values = []
                except mariadb.Error as e:
                    print(e)
                    self.conn.rollback()
                    break
            else:
                i += 1

    def get_api_data(self, api: str):
        """
        获取接口数据
        :param str api: 接口地址
        :return:
        """
        response = requests.get(api, proxies=self.proxies, verify=False)
        return response.json()


if __name__ == "__main__":
    o = IPTVSpider()
    o.run()
