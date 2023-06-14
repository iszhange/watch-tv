import { app, BrowserWindow, session, net, netLog, ipcMain } from "electron";
import path from "path";

const config = {
  apiHost: "http:localhost:81",
}

// 获取国家数据
async function netGetCountries(query: object) {
  const params = new URLSearchParams()
  Object.entries(query).forEach(([key, val]) => {
    params.append(key, val)
  })

  const url = new URL(config.apiHost + '/countries')
  url.search = params.toString()

  const request = net.request({
    method: "GET",
    url: url.toString(),
  })
  let data = ''
  request.on('response', (response) => {
    console.log(`STATUS: ${response.statusCode}`)
    console.log(`HEADERS: ${JSON.stringify(response.headers)}`)
    response.on('data', (chunk) => {
      console.log(`BODY: ${chunk}`)
      data += chunk
    })
    response.on('end', () => {
      console.log('No more data in response.')
      return data
    })
  })
  request.end()
}

const createWindow = () => {
  const win = new BrowserWindow({ 
    width: 800,
    webPreferences: {
      // contextIsolation: false, // 是否开启隔离上下文
      // nodeIntegration: true, // 渲染进程使用Node API
      // allowRunningInsecureContent: true, // 允许网站在HTTPS中加载或执行非安全源(HTTP) 中的脚本代码、CSS或插件
      preload: path.join(__dirname, "../dist-electron/preload.js"), // 需要引用js文件
    },
  });

  // 如果打包了，渲染index.html
  if (app.isPackaged) {
    win.loadFile(path.join(__dirname, "../index.html"));
  } else {
    let url = "http://localhost:3000"; // 本地启动的vue项目路径
    win.loadURL(url);
  }
};


app.whenReady().then(() => {
  // 禁用CSP
  session.defaultSession.webRequest.onHeadersReceived((details, callback) => {
    callback({
      responseHeaders: {
        ...details.responseHeaders,
        'Content-Security-Policy': '*',
      }
    })
  })

  // 网络请求事件
  ipcMain.handle('net:api:countries', (event, params) => netGetCountries(params))

  // 创建窗口
  createWindow(); 
  app.on("activate", () => {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

// 关闭窗口
app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit();
  }
});

