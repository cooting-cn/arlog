/*基础请求信息*/
const base = {
    /*    url: "https://arlog.cn",/!*默认路由, 生产要修改成当前域名*!/*/
    url: "http://127.0.0.1",/*默认路由, 生产要修改成当前域名*/
    login: "/api/v2/login", /*登录*/
    getCode: "/api/v2/get-code", /*查询gitee代码分析*/
    getArt: "/api/v1/get-art",/*查询文章*/
}


export default base
