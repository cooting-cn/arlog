<script setup>
import {h} from "vue"
import {NButton} from "naive-ui"
import api from "@/api/api.js";

/*数据加载变量*/
const loading = ref(false)
/*按钮触发函数*/
const sendMail = (rowData) => {
  $message.info(`send mail to ${rowData.id}`)
}


/*分页设置*/
const pagination = reactive({
  page: 2,  // 当前页
  pageSize: 12,  // 每页显示的记录数
  pageCount: 8,  // 总页数
  prefix({itemCount}) {  // 自定义前缀文本
    return `总数 ${itemCount}`  // 显示 "Total is {itemCount}."
  }
})

/*自动触发*/
onMounted(() => {
  api.getArt(params).then(res => {

    data.value = res.data.result.arts
    pagination.itemCount = res.data.result.total
    console.log(data.value)
  })

})

// 查询参数
const params = reactive({
  page: 1,
})

/*触发分页*/
function page(currentPage) {

  if (!loading.value) {
    loading.value = true
    /*显示的当前页*/
    pagination.page = currentPage
    /*请求的后端的参数页*/
    params.value.page = currentPage
    /*请求文章*/
    api.getArt(params).then(res => {
      data.value = res.data.result.arts
      console.log(data.value)
    })

    loading.value = false  // 请求完成后取消加载状态
  }
}

/*导航分类设置*/
const columns = ref([
  {
    title: "ID",
    key: "id",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center'  // 表头居中对齐
  },
  {
    title: "文章",
    key: "title",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center'  // 表头居中对齐

  },
  {
    title: "描述",
    key: "desc",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "分类",
    key: "sortId",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "标签",
    key: "tags",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "发布",
    key: "open",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center'  // 表头居中对齐
  },
  {
    title: "创建时间",
    key: "created_at",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "更新时间",
    key: "updated_at",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "Action",
    key: "actions",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
    render(row) {
      return h(
          NButton,
          {
            size: "small",
            onClick: () => sendMail(row)
          },
          {default: () => "Send Email"}
      )
    }
  }
])
/*列表数据*/
const data = ref([])


</script>

<template>


  <n-card class="h-100% ">
    <!--  筛选 添加框  -->
    <n-card class="h-80px " style="background-color: rgba(250, 250, 252, 1)">
      <n-space class="items-center">
        <div class="w-300px">
          <n-space class="items-center">
            <span>文章</span>
            <n-input placeholder="文章" type="text"/>
          </n-space>

        </div>

        <div class="w-300px">
          <n-space class="items-center">
            <span>分类</span>
            <n-input placeholder="分类" type="text"/>
          </n-space>

        </div>


        <n-button type="tertiary">
          搜索
          <template #icon>
            <n-icon
                :class="'i-ep-search'"
                size="15px"
            />
          </template>
        </n-button>
        <n-button type="primary">
          新增
          <template #icon>
            <n-icon
                :class="'i-ep-plus'"
                size="15px"
            />
          </template>
        </n-button>
      </n-space>

    </n-card>

    <!-- 数据表单   -->
    <n-data-table
        :columns=columns
        :data=data
        :loading=loading
        :pagination=pagination
        class="mt-20"
        min-height="690"
        remote
        size="large"
        @update:page="page"

    />

  </n-card>

</template>

<style scoped>

</style>