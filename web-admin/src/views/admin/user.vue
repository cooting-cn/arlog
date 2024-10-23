<script setup>

import {h} from "vue"
import {NButton, NPopconfirm, NAvatar, NTag} from "naive-ui"
import api from "@/api/api.js"


/*数据加载变量*/
// 控制分页的显示
const loading = ref(false)


/*分页设置*/
const pagination = reactive({
  page: 1,  // 当前页
  pageSize: 12,  // 每页显示的记录数
  pageCount: null,  // 总页数
  prefix({itemCount}) {  // 自定义前缀文本
    return `总数 ${itemCount}`  // 显示 "Total is {itemCount}."
  }
})

/*自动触发*/
onMounted(() => {
  /*加载数据中*/
  loading.value = true

  $loadingBar.start()
  /*查询用户*/
  api.getAllUsers(params).then(
      res => {
        // 获取文章用户
        data.value = res.data.result.tags
        /*获取数据条数*/
        pagination.itemCount = res.data.result.total
        /*关闭加载*/
        $loadingBar.finish()
        loading.value = false
      }
  )


})

// 查询参数
const params = reactive({
  page: 1,
})

/*触发分页*/
function page(currentPage) {

  if (!loading.value) {
    $loadingBar.start()
    loading.value = true
    /*显示的当前页*/
    pagination.page = currentPage
    /*请求的后端的参数页*/

    params.page = currentPage

    /*请求用户*/
    api.getTags(params).then(res => {
      data.value = res.data.result.tags
      /*关闭加载*/
      $loadingBar.finish()
      loading.value = false
    })


  }
}


/*导航用户设置*/
const columns = ref([
  {
    title: "头像",
    key: "a-img",
    width: '5%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
    render(row) {
      return h(
          NAvatar,
          {
            size: 'medium',
            src: row['a-img']
          },
      )
    }
  },

  {
    title: "用户",
    key: "username",
    width: '5%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐


  },
  {
    title: "权限",
    key: "power",
    width: '5%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
    render(row) {
      return h(
          NTag,
          {type: row.power === 'admin' ? 'warning' : 'info', bordered: false}, // 你可以根据需要调整 type
          {default: () => row.power === 'admin' ? '管理员' : '游客'} // 标签内容
      )
    }

  },
  {
    title: "微信",
    key: "wechat",
    width: '5%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐


  },
  {
    title: "令牌otp",
    key: "otp",
    width: '5%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
    ellipsis: {
      tooltip: true //带提示的省略
    },
    render(row) {
      return h(
          NButton, {
            disabled: row.otp.length > 0,
            size: "small",
            onClick: () => getOtp(row.username)
          }, {default: () => "开启二级验证"}
      )
    }
  },
  {
    title: "登录ip",
    key: "ip",
    width: '5%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "创建时间",
    key: "created_at",
    width: '10%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "更新时间",
    key: "updated_at",
    width: '10%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "功能",
    key: "actions",
    width: '5%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
    render(row) {
      return h(
          NPopconfirm,
          {
            onPositiveClick: () => del(row), // 确认后调用的函数
            onNegativeClick: () => $message.info('取消删除')  // 取消后的操作
          },
          {
            trigger: () => h(NButton, {disabled: row.power === 'admin', size: "small"}, {default: () => "删除"}),
            default: () => "是否删除用户？" // 弹出确认框的内容
          }
      )
    }
  }
])
/*列表数据*/
const data = ref([])


/*新增*/
const showModal = ref(false)

function Add() {
  showModal.value = true
}


const formValue = ref(
    {
      name: ""
    })

/*确定添加按钮*/
function add() {
  api.addTag(formValue.value).then(
      res => {
        // 关闭弹窗
        showModal.value = false
        // 重新当前页面数据
        page(pagination.page)
        // 提示成功
        $message.success("添加成功" + res.data.msg)
        // 重置表单
        formValue.value = {
          name: ""
        }
      }
  )
}

/*删除按钮*/
const del = (rowData) => {
  //赋值给动态表单
  formValue.value.name = rowData.name

  api.deleteTag(formValue.value).then(
      res => {
        console.log(formValue.value)
        // 删除成功
        $message.success("删除成功" + res.data.msg)
        // 重新当前页面数据
        page(pagination.page)
      }
  )
}

/*绑定otp*/


const showModalOtp = ref(false)
const otp = ref({
  username: ""
})
const otpTxt = ref('')

const bindotp = ref({
  username: "",
  otp: "",
  coding: ""
})

/*获取opt编码*/
function getOtp(username) {

  otp.value.username = username
  /*获取绑定*/
  api.getOtp(otp.value).then(
      res => {
        otpTxt.value = res.data.result.url
        // 提示成功
        $message.success("获取" + res.data.msg)
        /*绑定数据*/
        bindotp.value.username = username
        bindotp.value.otp = res.data.result.secret
      }
  )

  //开启弹窗
  showModalOtp.value = true
}

/*绑定otp*/
function bindOtp() {

  api.bindOtp(bindotp.value).then(
      res => {
        $message.success("二级验证绑定已开启" + res.data.msg)
      }
  )
  //关闭弹窗
  showModalOtp.value = false
}

</script>

<template>
  <n-card class="h-100% min-w-900 ">
    <!--  筛选 添加框  -->
    <n-card class="flex " style="background-color: rgba(250, 250, 252, 1)">
      <div class="flex items-center w-100% ">
        <div class="w-500px ">
          <n-space class="items-center ">
            <span>用户</span>
            <n-input :input-props="{ autocomplete: 'wz' ,id: 'wz'}" disabled
                     placeholder="用户"
                     style="min-width: 250px"
                     type="text"/>
            <n-button disabled type="tertiary">
              搜索
              <template #icon>
                <n-icon
                    :class="'i-ep-search'"
                    size="15px"
                />
              </template>
            </n-button>
          </n-space>

        </div>


        <div class="absolute right-0 mr-20">

          <n-button class="mr-16" type="primary" @click="Add">
            新增
            <template #icon>
              <n-icon
                  :class="'i-ep-select'"
                  size="15px"
              />
            </template>
          </n-button>


        </div>

      </div>

    </n-card>

    <!-- 数据表单   -->
    <n-data-table
        :columns=columns
        :data=data
        :loading=loading
        :pagination="!loading ? pagination : false"
        class="mt-20 min-h-726"
        remote
        size="large"
        @update:page="page"

    />
  </n-card>

  <!--动态弹出框新增-->
  <n-modal v-model:show="showModal">
    <n-card
        :bordered="false"
        aria-modal="true"
        class="w-20% min-w-300"
        role="dialog"
        size="huge"
        title="添加用户"
    >
      <n-form :model="formValue" label-placement="left" label-width="auto">
        <n-form-item label="用户名">
          <n-input v-model:value="formValue.name" placeholder="用户名"/>
        </n-form-item>


      </n-form>


      <div class="grid place-items-end ">
        <n-button class="w-100" type="info" @click="add">
          添加
        </n-button>
      </div>

    </n-card>
  </n-modal>

  <!--动态弹出框 otp-->
  <n-modal v-model:show="showModalOtp">
    <n-card
        :bordered="false"
        aria-modal="true"
        class="w-20% min-w-300 "
        role="dialog"
        size="huge"
        title="绑定top"
    >


      <n-qr-code :size=300 :value=otpTxt style="padding: 0"/>


      <div class="text-center">
        <!--动态码-->
        <n-form-item>
          <n-input v-model:value="bindotp.coding" placeholder="谷歌动态码"/>
        </n-form-item>
        <!--开始绑定-->
        <n-button class="w-100" type="info" @click="bindOtp">
          确定绑定
        </n-button>
      </div>


    </n-card>
  </n-modal>
</template>

<style scoped>

</style>