// Modules to control application life and create native browser window
const {app, BrowserWindow,Menu, ipcMain,net} = require('electron');
const path = require('path');
const axios = require('axios').default;

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let mainWindow
let addwindow;
var list = [];


function createWindow () {
  // Create the browser window.
  mainWindow = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      nodeIntegration:true,
      preload: path.join(__dirname, 'preload.js')
    }
  });

  // and load the index.html of the app.
  mainWindow.loadFile('index.html')


  // Open the DevTools.
  // mainWindow.webContents.openDevTools()

  // Emitted when the window is closed.
  mainWindow.on('closed', function () {
    app.quit();
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    mainWindow = null
  })
 
  const mainMenu = Menu.buildFromTemplate(mainMenuTemplate);

  Menu.setApplicationMenu(mainMenu)
}
function createAddWindow(){
  addwindow = new BrowserWindow({
    width: 200,
    height: 300,
    title: 'add event list window',
    webPreferences: {
      nodeIntegration:true,
      preload: path.join(__dirname, 'preload.js')
    }
  });
  // and load the index.html of the app.
  addwindow.loadFile('addwindow.html')


//garbage collect addwindow.
  addwindow.on('close', function(){
    addwindow = null;
  });
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', createWindow)
// Quit when all windows are closed.
app.on('window-all-closed', function () {
  // On macOS it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== 'darwin') app.quit()
})

app.on('activate', function () {
  // On macOS it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (mainWindow === null) createWindow()
})

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.

//Catch item:add.
ipcMain.on('item:add', function(e){
  addwindow.loadFile('addwindow.html');
  //mainWindow.webContents.send('item:add', item);
  //addwindow.close(); 
});

ipcMain.on('event:add', function(e,item){
  console.log(item);
  addwindow.loadFile('addwindow.html');
  addwindow.webContents.send('event:add', item);
  axios.post('http://localhost:8080/event',item)
  .then(function(response){
    console.log(response.status);
  }).catch(function(err){
    console.log(err);
  })
  addwindow.close(); 
});

//create menu template

const mainMenuTemplate = [
  {
    label : 'File',
    submenu:[
      {
        label: 'Add Event',
        click(){
          createAddWindow();
        }
      },
      {
        label : 'Refresh',
        click(){
         axios.get('http://localhost:8080/events')
         .then(function(response){
           var r = JSON.stringify(response.data).split(",{");
          for(var i = 0; i < r.length;i++){
            var temp = "{"+r[i];
            if(i == 0){
              console.log("before slice first " + temp);
             temp = r[i].substr(1);
             console.log("after slice first " + temp);
            }
            else if(i==r.length-1){
              console.log("before slice last " + temp);
              temp = r[i].slice(0,r[i].length-1);
              temp = "{"+temp;
            }
            if(!list.includes(temp)){
              list.push(temp);
              mainWindow.webContents.send('item:add',temp);
            }
          }
         }).catch(function(err){
           console.log(err);
         })
           
          }
          
          
        
      },
      {
        label : 'Quit',
        click(){
          app.quit();
        }
      }      
    ]
  },
{
  label: 'help'
}
]
//If platform is MAC add an empty object to the start of the template. 
if(process.platform == 'darwin'){
  mainMenuTemplate.unshift({})
}
if(process.env.NODE_ENV !== 'production'){
  mainMenuTemplate.push({
    label : 'Developer Tools',
    submenu:[
      {
        label: 'Toggle DevTools',
        click(item,focusedWindow){
          focusedWindow.toggleDevTools();
        }
      },
      {
        role: 'reload'
      }
    ]
  });
}