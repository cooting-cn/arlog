import axios from "axios";
import base from "@/api/base.js";
import blogStore from "@/stores/arlog.js";
import {useMessage} from "naive-ui";
import {useRouter} from "vue-router";

const message = useMessage();
const router = useRouter();
const st = blogStore();

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
);


// 响应拦截器：统一处理错误响应
req.interceptors.response.use(
    response => {
        const res = response.data;

        // 成功响应的处理
        if (res.code === 200) {
            return res; // 返回响应数据
        } else {
            // 统一处理 code 不为 200 的情况
            message.error(res.msg || '请求失败');
            return Promise.reject(new Error(res.msg || 'Error'));
        }
    },
    error => {
        // 错误响应的处理
        if (error.response) {
            const status = error.response.status;

            switch (status) {
                case 401:
                    // 如果是未授权（如 token 过期）
                    message.error('登录已过期，请重新登录');
                    st.token = ''; // 清除本地 token
                    router.push({name: 'login'}); // 跳转到登录页面
                    break;
                case 403:
                    message.error('您没有权限执行此操作');
                    break;
                case 500:
                    message.error('服务器内部错误，请稍后再试');
                    break;
                default:
                    message.error('发生未知错误');
            }
        } else {
            if (error.message.includes('timeout')) {
                message.error('请求超时，请稍后重试');
            } else {
                message.error('网络错误，请检查您的网络连接');
            }
        }
        return Promise.reject(error);
    }
);

export default req