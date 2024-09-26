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
    }
}


export default api