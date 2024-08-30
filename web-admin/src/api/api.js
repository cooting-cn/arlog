/*请求*/
import req from "@/utils/req.js";
import base from "@/api/base.js";

const api = {
    /*登录请求*/
    postLogin(data) {
        return req.post(base.login, data)
    },
}

export default api