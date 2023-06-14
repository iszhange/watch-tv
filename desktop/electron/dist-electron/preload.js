"use strict";const e=require("electron");e.contextBridge.exposeInMainWorld("WatchTVAPI",{countries:n=>e.ipcRenderer.invoke("net:api:countries",n)});
