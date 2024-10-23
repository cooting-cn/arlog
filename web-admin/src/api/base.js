/*基础请求信息*/
const base = {

    url: baseUrl,/*默认路由, 生产要修改成当前域名*/
    /*登录*/
    login: "/api/v2/login",
    /*查询gitee代码分析*/
    getCode: "/api/v2/get-code",
    /*查询文章*/
    getArt: "/api/v1/get-art",
    /*搜索文章*/
    searchArt: "/api/v1/sear-art",
    /*批量删除文章*/
    deleteArt: "/api/v1/del-art",
    /*修改文章*/
    upArt: "/api/v1/up-art",
    /*添加文章*/
    addArt: "/api/v1/add-art",
    /*上传图片*/
    uploadImage: "/api/v1/upload-image",
    /*获取标签*/
    getTags: "/api/v1/get-tag",
    /*添加标签*/
    addTag: "/api/v1/add-tag",
    /*删除标签*/
    deleteTag: "/api/v1/del-tag",
    /*获取文章分类*/
    getArtCats: "/api/v1/get-art-cats",
    /*添加文章分类*/
    addArtCat: "/api/v1/add-art-cat",
    /*查询分类*/
    getSort: "/api/v1/get-sort",
    /*添加分类*/
    addSort: "/api/v1/add-sort",
    /*删除分类*/
    deleteSort: "/api/v1/del-sort",
    /*查询用户*/
    getAllUser: "/api/v1/get-all-user",
    /*获取otp*/
    getOtp: "/api/v1/get-otp",
    /*绑定otp*/
    bindOtp: "/api/v1/bind-otp",
}


export default base
