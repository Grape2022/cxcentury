脚本存放路径  
C:\Users\Administrator\AppData\Roaming\apifox\ExternalPrograms\com\cxcentury\

---

文件案例,文件路径
C:\Users\Administrator\AppData\Roaming\apifox\ExternalPrograms\com\cxcentury\model\productBase\FileUploadByVerify.go

实现代码
  // 获取图片上传信息:       C:\Users\Administrator\AppData\Roaming\apifox\ExternalPrograms\com\cxcentury\model\productBase
    let goResult = pm.execute('com/cxcentury/model/productBase/FileUploadByVerify.go', [1, 2], {
        className: "main",
        model: "Test",
        paramTypes: ['int', 'int'],
        windowsEncoding: 'utf-8',
    })
    goResult = JSON.parse(goResult)

------
