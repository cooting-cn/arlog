/*请求*/
import req from "@/utils/req.js";
import base from "@/api/base.js";

const api = {
    /*登录请求*/
    postLogin(data) {
        return req.post(base.login, data)
    },
    /*查询gitee代码分析*/
    getCode() {
        return req.get(base.getCode)
    },
    /*查询文章*/
    getArt(data) {
        return req.get(base.getArt, {params: data})
    },
    /*搜索文章*/
    searchArt(data) {
        return req.post(base.searchArt, data)
    },
    /*批量删除文章*/
    deleteArt(data) {
        return req.post(base.deleteArt, data)
    },
    /*添加文章*/
    addArt(data) {
        return req.post(base.addArt, data)
    },
    /*编辑文章*/
    upArt(data) {
        return req.post(base.upArt, data)
    },
    /*上传图片*/
    uploadImage(data) {
        return req.post(base.uploadImage, data)
    },
    /*查询分类*/
    getSort(data) {
        return req.get(base.getSort, {params: data})
    },
    /*删除分类*/
    deleteSort(data) {
        return req.post(base.deleteSort, data)
    },
    /*添加分类*/
    addSort(data) {
        return req.post(base.addSort, data)
    },
    /*获取标签*/
    getTags(data) {
        return req.get(base.getTags, {params: data})
    },
    /*添加标签*/
    addTag(data) {
        return req.post(base.addTag, data)
    },
    /*删除标签*/
    deleteTag(data) {
        return req.post(base.deleteTag, data)
    },
    /*查询所有用户*/
    getAllUsers(data) {
        return req.get(base.getAllUser, {params: data})
    },
    /*获取otp*/
    getOtp(data) {
        return req.get(base.getOtp, {params: data})
    },
    /*绑定otp*/
    bindOtp(data) {
        return req.post(base.bindOtp, data)
    },
    /*otp登录*/
    otp(data) {
        return req.post(base.otp, data)
    },
}


export default api