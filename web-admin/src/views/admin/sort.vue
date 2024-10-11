<script setup>

import {h} from "vue"
import {NButton, NPopconfirm} from "naive-ui"
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
  /*查询分类*/
  api.getSort(params).then(
      res => {
        // 获取文章分类
        data.value = res.data.result.sorts
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

    /*请求分类*/
    api.getSort(params).then(res => {
      data.value = res.data.result.sorts
      /*关闭加载*/
      $loadingBar.finish()
      loading.value = false
    })


  }
}


/*导航分类设置*/
const columns = ref([
  {
    title: "ID",
    key: "id",
    width: '5%',
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center'  // 表头居中对齐
  },

  {
    title: "分类",
    key: "name",
    width: '20%',
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
            onPositiveClick: () => deSrt(row), // 确认后调用的函数
            onNegativeClick: () => $message.info('取消删除')  // 取消后的操作
          },
          {
            trigger: () => h(NButton, {size: "small"}, {default: () => "删除"}),
            default: () => "是否删除分类？" // 弹出确认框的内容
          }
      )
    }
  }
])
/*列表数据*/
const data = ref([])


/*新增*/
const showModal = ref(false)

function AddArt() {
  showModal.value = true
}


const formValue = ref(
    {
      name: ""
    })

/*确定添加按钮*/
function addArt() {
  api.addSort(formValue.value).then(
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
const deSrt = (rowData) => {
  //赋值给动态表单
  formValue.value.name = rowData.name

  api.deleteSort(formValue.value).then(
      res => {
        console.log(formValue.value)
        // 删除成功
        $message.success("删除成功" + res.data.msg)
        // 重新当前页面数据
        page(pagination.page)
      }
  )
}


</script>

<template>
  <n-card class="h-100% min-w-900 ">
    <!--  筛选 添加框  -->
    <n-card class="flex " style="background-color: rgba(250, 250, 252, 1)">
      <div class="flex items-center w-100% ">
        <div class="w-500px ">
          <n-space class="items-center ">
            <span>分类</span>
            <n-input :input-props="{ autocomplete: 'wz' ,id: 'wz'}" disabled
                     placeholder="分类"
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

          <n-button class="mr-16" type="primary" @click="AddArt">
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

  <!--动态弹出框-->
  <n-modal v-model:show="showModal">
    <n-card
        :bordered="false"
        aria-modal="true"
        class="w-20% min-w-300"
        role="dialog"
        size="huge"
        title="添加分类"
    >
      <n-form :model="formValue" label-placement="left" label-width="auto">
        <n-form-item label="分类名">
          <n-input v-model:value="formValue.name" placeholder="分类名"/>
        </n-form-item>


      </n-form>


      <div class="grid place-items-end ">
        <n-button class="w-100" type="info" @click="addArt">
          添加
        </n-button>
      </div>

    </n-card>
  </n-modal>
</template>

<style scoped>

</style>