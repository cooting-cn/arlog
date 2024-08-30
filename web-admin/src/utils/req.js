import axios from "axios";
import base from "@/api/base.js";


/*创建请求对象*/
const req = axios.create(
    {
        baseURL: base.url,
        /*线上修改成请求当前域名*/

        timeout: 5000,
    }
)
/*注册请求拦截*/

/*所有post请求添加成json格式*/
req.interceptors.request.use((config) => {
        if (config.method === 'POST') {
            config.headers.append('Content-Type', 'application/json');
        }
        return config;
    },
    error => Promise.reject(error)
)
export default req