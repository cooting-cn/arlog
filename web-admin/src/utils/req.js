import axios from "axios";
import base from "@/api/base.js";
import blogStore from "@/stores/arlog.js";


const st = blogStore()

/*创建请求对象*/
const req = axios.create(
    {
        baseURL: base.url,
        /*线上修改成请求当前域名*/

        timeout: 5000,
    }
)
/*注册请求拦截*/

/*添加请求token*/
/*所有post请求添加成json格式*/
req.interceptors.request.use((config) => {

        config.headers['token'] = st.token

        if (config.method === 'POST') {
            config.headers.append('Content-Type', 'application/json');
        }
        return config
    },
    error => Promise.reject(error)
);


// 响应拦截器：统一处理错误响应
// 添加响应拦截器
req.interceptors.response.use(res => {

        // 对响应数据做点什么
        const data = res.data
        // 访问成功 清理 之前的请求
        $message.destroy("loadingMessage")
        switch (data.code) {

            case 203:
                // 显示消息
                // 无论成功还是失败，都会执行

                window.$message.warning(data.msg)
                break
            case 204:
                // 显示消息
                window.$message.warning(data.msg)
                break
            case 205:
                // 显示消息
                window.$message.warning(data.msg)
                break

        }
        return res
    }, error => Promise.reject(error)
);


export default req