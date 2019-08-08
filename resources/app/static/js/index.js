let index = {
    about: function(html) {
        let c = document.createElement("div");
        c.innerHTML = html;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();
        })

        // ---------- 设置默认值
        txtPlain =document.getElementById("txtPlain")
        txtKey =document.getElementById("txtKey")
        txtPlain.value = "11111111111111111111111111111111"
        txtKey.value = "11111111111111111111111111111111"
    },
    listen: function() {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    index.about(message.payload);
                    return {payload: "payload"};
                    break;
            case "index_out_test":
                asticode.notifier.info(message.payload);
                break;
            }
        });
    },
    btn_crypt: function(){
        // 为页面控件定义变量
        txtPlain =document.getElementById("txtPlain")
        txtKey =document.getElementById("txtKey")
        txtCipher =document.getElementById("txtCipher")

        var plain = txtPlain.value;
        var key = txtKey.value;
        s = "plain:" + plain + ", key:" + key 
        console.log(s)
//        alert(s)


        // 清空输入框
        txtCipher.value = ""

        let params = {"plain":plain,"key":key}
        console.log(params)
        // Create message
        let message = {"name": "btn_crypt","payload":params};

        astilectron.sendMessage(message, function(message) {
            console.log("--> Go语言执行完毕收到回调, 收到: " + message)
            console.log(message)

            var ret = message.payload
            if(ret.ret == false){
                alert(ret.message)
                return 
            }
            txtCipher.value =ret.value
        })
    }
};
