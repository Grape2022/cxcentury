try {
    const request = pm.request
    let token = pm.globals.get("token")

    query = pm.request.url.query
    let companyName = pm.variables.get("companyName")

    let rag = new RegExp(companyName, "g")
    let funcDateTime = (num = 0) => {
        nums = new Date().getTime() + (1000 * 60 * 60 * 8) + (1000 * 60 * 60 * 24 * num)
        let DateTime = new Date(nums)
        let date = ""
        for (let i = 1; i < 11; i++) {
            if (JSON.stringify(DateTime)[i] != undefined) {
                date += JSON.stringify(DateTime)[i]
            }
        }
        let time = ""
        for (let i = 12; i < 20; i++) {
            if (JSON.stringify(DateTime)[i] != undefined) {
                time += JSON.stringify(DateTime)[i]
            }
        }
        let dateTime = date + " " + time
        let timeList = [date, time, dateTime]
        return timeList
    }
    timeList = funcDateTime()

    //获取logisticsCompanyId
    const queryCompanyId = {
        url: `http://192.168.10.101:8182/prod-api/logisticsCompanyForgroups/list?companyNames=&isAsc=desc&orderByColumn=createTime`,
        method: `GET`,
        header: {
            "Authorization": token,
        }
    };
    pm.sendRequest(queryCompanyId, function (err, res) {

        for (let i = 0; i < res.json().data.length; i++) {
            if (companyName != undefined && rag.test(res.json().data[i].companyName)) {

                //获取渠道,获取 destinationCountry ,仓库代码
                const showChannelName = {
                    url: `http://192.168.10.101:8182/prod-api/logisticsChannelForgroups/list?logisticsCompanyId=` + res.json().data[i].logisticsCompanyId + `&channelNameStr=`,
                    method: "GET",
                    header: {
                        "Authorization": token,
                    },
                }
                pm.sendRequest(showChannelName, function (err_2, res_2) {

                    let maxLogisticsChannelForgroupId = 0;
                    let isNum = 0
                    let i = 0;
                    while (i < res_2.json().data.length) {

                        if (maxLogisticsChannelForgroupId < res_2.json().data[i].logisticsChannelForgroupId) {
                            maxLogisticsChannelForgroupId = res_2.json().data[i].logisticsChannelForgroupId
                            isNum = i
                        }
                        i++
                    }

                    console.log(res_2.json().data[isNum].channelName)

                    //获取fbaCode
                    const showFbaCode = {
                        url: `http://192.168.10.101:8182/prod-api/logisticsAreass/subAreas?areaId=` + res_2.json().data[isNum].destinationCountry,
                        method: "GET",
                        header: {
                            "Authorization": token,
                        },
                    }
                    pm.sendRequest(showFbaCode, function (err_4, res_3) {

                        //添加报价
                        //	http://192.168.10.101:8182/prod-api/logisticsChannelForgroupDates

                        let numStr = Math.trunc(Math.random() * 100 + 10)

                        const afterAddLogisticsChannelForgroupDates = {
                            url: `http://192.168.10.101:8182/prod-api/logisticsChannelForgroupDates`,
                            method: "POST",
                            header: {
                                "Authorization": token,
                                "Content-Type": "application/json",
                            },
                            body: {
                                mode: "raw",
                                raw: JSON.stringify({
                                    "filePath": "",
                                    "logisticsChannelForgroup": {
                                        "channelSheetName": ""
                                    },
                                    "logisticsChannelForgroupDate": {
                                        "effectTime": timeList[0],
                                        "logisticsChannelForgroupId": maxLogisticsChannelForgroupId,
                                        "channelSheetName": ""
                                    },
                                    "logisticsChannelForgroupPartBo": {
                                        "logisticsChannelForgroupPart": {},
                                        "logisticsChannelForgroupPartQuotationList": [
                                            {
                                                "startNumber": 0,
                                                "endNumber": 50
                                            },
                                            {
                                                "startNumber": 50.01,
                                                "endNumber": 100
                                            },
                                            {
                                                "startNumber": 100.01,
                                                "endNumber": 1000
                                            }
                                        ]
                                    },
                                    "logisticsChannelForgroupSetBoList": [
                                        {
                                            "logisticsChannelForgroupSet": {
                                                "sendDay": 0,
                                                "setType": 1
                                            },
                                            "logisticsChannelForgroupDatePriceBoList": [
                                                {
                                                    "orderNo": 0,
                                                    "expensesPrice": 0,
                                                    "startNumber": 0,
                                                    "endNumber": 50
                                                },
                                                {
                                                    "orderNo": 0,
                                                    "expensesPrice": 0,
                                                    "startNumber": 50.01,
                                                    "endNumber": 100
                                                },
                                                {
                                                    "orderNo": 0,
                                                    "expensesPrice": 0,
                                                    "startNumber": 100.01,
                                                    "endNumber": 1000
                                                }
                                            ],
                                            "logisticsChannelForgroupDatePositionBoList": [],
                                            "logisticsChannelForgroupSetElementList": [
                                                {
                                                    "areaId": null,
                                                    "fbaCode": "AMD1",
                                                    "warehouseId": null,
                                                    "zip": ""
                                                },]
                                        }
                                    ]
                                })
                            },
                        }

                        afterAddLogisticsChannelForgroupDates.body.raw = JSON.parse(afterAddLogisticsChannelForgroupDates.body.raw)
                        afterAddLogisticsChannelForgroupDates.body.raw.logisticsChannelForgroupSetBoList[0].logisticsChannelForgroupSetElementList = []

                        // console.log(`res_3`, res_3.json().data)

                        console.log(`res_3.json()`, res_3.json().data[0].name)
                        afterAddLogisticsChannelForgroupDates.body.raw.logisticsChannelForgroupSetBoList=[]

                        if (`res_3.json()`, res_3.json().data[0].name == "美国") {
                            console.log(`res_3.json()`, res_3.json().data[0].children.length)
                            for (let x = 0; x < res_3.json().data[0].children.length; x++) {

                                console.log(`children[x]`, res_3.json().data[0].children[x].fbaCodes.length)

                                for (let y = 0; y < res_3.json().data[0].children[x].fbaCodes.length; y++) {



                                    // afterAddLogisticsChannelForgroupDates.body.raw.logisticsChannelForgroupSetBoList=[
                                    //     ...afterAddLogisticsChannelForgroupDates.body.raw.logisticsChannelForgroupSetBoList,
                                    //     {
                                    //         "logisticsChannelForgroupSet": {
                                    //             "sendDay": Math.trunc(Math.random() * 30 + 12),
                                    //             "setType": 1
                                    //         },
                                    //         "logisticsChannelForgroupDatePriceBoList": [
                                    //             {
                                    //                 "orderNo": 0,
                                    //                 "expensesPrice": String(numStr + 12),
                                    //                 "startNumber": 0,
                                    //                 "endNumber": 50
                                    //             },
                                    //             {
                                    //                 "orderNo": 0,
                                    //                 "expensesPrice": String(numStr + 5),
                                    //                 "startNumber": 50.01,
                                    //                 "endNumber": 100
                                    //             },
                                    //             {
                                    //                 "orderNo": 0,
                                    //                 "expensesPrice": String(numStr),
                                    //                 "startNumber": 100.01,
                                    //                 "endNumber": 1000
                                    //             }
                                    //         ],
                                    //         "logisticsChannelForgroupDatePositionBoList": [],
                                    //         "logisticsChannelForgroupSetElementList": [
                                    //             {
                                    //                 "areaId": null,
                                    //                 "fbaCode":res_3.json().data[0].children[x].fbaCodes[y].fbaCode,
                                    //                 "warehouseId": null,
                                    //                 "zip": ""
                                    //             },]
                                    //     }
                                    // ]


                                }
                            }

                        } else {
                            console.log(res_3.json().data)
                            for (let x = 0; x < res_3.json().data.length; x++) {
                                console.log(res_3.json().data[x].fbaCodes)
                                for (let y = 0; y < res_3.json().data[x].fbaCodes.length; y++) {


                                    // afterAddLogisticsChannelForgroupDates.body.raw.logisticsChannelForgroupSetBoList=[
                                    //     ...afterAddLogisticsChannelForgroupDates.body.raw.logisticsChannelForgroupSetBoList,
                                    //     {
                                    //         "logisticsChannelForgroupSet": {
                                    //             "sendDay": Math.trunc(Math.random() * 30 + 12),
                                    //             "setType": 1
                                    //         },
                                    //         "logisticsChannelForgroupDatePriceBoList": [
                                    //             {
                                    //                 "orderNo": 0,
                                    //                 "expensesPrice": String(numStr + 12),
                                    //                 "startNumber": 0,
                                    //                 "endNumber": 50
                                    //             },
                                    //             {
                                    //                 "orderNo": 0,
                                    //                 "expensesPrice": String(numStr + 5),
                                    //                 "startNumber": 50.01,
                                    //                 "endNumber": 100
                                    //             },
                                    //             {
                                    //                 "orderNo": 0,
                                    //                 "expensesPrice": String(numStr),
                                    //                 "startNumber": 100.01,
                                    //                 "endNumber": 1000
                                    //             }
                                    //         ],
                                    //         "logisticsChannelForgroupDatePositionBoList": [],
                                    //         "logisticsChannelForgroupSetElementList": [
                                    //             {
                                    //                 "areaId": null,
                                    //                 "fbaCode": res_3.json().data[x].fbaCodes[y].fbaCode,
                                    //                 "warehouseId": null,
                                    //                 "zip": ""
                                    //             },]
                                    //     }
                                    // ]

                                }
                            }
                        }
                        afterAddLogisticsChannelForgroupDates.body.raw = JSON.stringify(afterAddLogisticsChannelForgroupDates.body.raw)

                        pm.sendRequest(afterAddLogisticsChannelForgroupDates, function (err_5, res_4) {
                            console.log(res_4.json())
                        })

                    })
                })

            };
        };
    })
} catch (error) {
    console.log("添加渠道/添加仓库报价-后置脚本:", error)
}