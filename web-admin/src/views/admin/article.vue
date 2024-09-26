<script setup>
import {h} from "vue"
import {NButton, NSwitch} from "naive-ui"
import api from "@/api/api.js";
import {MdEditor} from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'

/*数据加载变量*/
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
        songs.value = res.data.result.sorts;

        // 将分类数据转换为映射对象
        categoryMap.value = songs.value.reduce((acc, cur) => {
          acc[cur.id] = cur.name;
          return acc;
        }, {})

      }
  )

  /*获取数据*/
  api.getArt(params).then(res => {
    /*获取数据*/
    data.value = res.data.result.arts
    /*获取数据条数*/
    pagination.itemCount = res.data.result.total
    /*关闭加载*/
    $loadingBar.finish()
    loading.value = false
  })

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

    /*请求文章*/
    api.getArt(params).then(res => {
      data.value = res.data.result.arts
      /*关闭加载*/
      $loadingBar.finish()
      loading.value = false
    })


  }
}


// 自定义轨道样式
function railStyle({checked}) {
  const style = {};
  if (checked) {
    style.background = "rgba(190,150,229,0.8)";

  }
  return style;
}


// 创建一个映射对象用于快速查找分类名称
const categoryMap = ref({})

/*导航分类设置*/
const columns = ref([
  {
    type: "selection"
  },
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
    render(row) {
      return categoryMap.value[row.sortId] || '未知'
    }
  },
  {
    title: "发布",
    key: "open",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
    render(row) {
      return h(
          NSwitch,
          {
            checkedValue: 1,
            uncheckedValue: 0,
            value: row.open,
            size: "medium",
            railStyle: railStyle,
            onClick: () => handleOpenChange(row)
          }
      )
    }
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
    title: "功能",
    key: "actions",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
    render(row) {
      return h(
          NButton,
          {
            size: "small",
            onClick: () => etArt(row)
          },
          {default: () => "编辑"}
      )
    }
  }
])
/*列表数据*/
const data = ref([])
/*选中行*/

// 响应式数据
const checkedRowKeys = ref([])

// 行的唯一标识
const rowKey = (row) => row.id
// 处理行选中
// 处理行选中事件
const handleCheck = (rowKeys) => {
  checkedRowKeys.value = rowKeys


}

/*删除文章按钮*/
const delArt = reactive({
  id: []
})

function DeleteArticle() {

  if (checkedRowKeys.value.length > 0) {
    delArt.id = checkedRowKeys.value
    api.deleteArt(delArt).then(
        res => {
          // 删除成功
          $message.success("删除成功" + res.data.msg)
          // 重新当前页面数据
          page(pagination.page)
        }
    )


    return
  }

  $message.error("没有删除的数据" + checkedRowKeys.value.length)

}

/*搜索文章*/
const seArt = reactive({
  title: ""
})

function SearchArticle() {
  // 获取搜索框中的值
  $loadingBar.start()
  api.searchArt(seArt).then(
      res => {
        //获取搜索的文章
        data.value = res.data.result.arts
        /*获取数据条数*/
        pagination.itemCount = res.data.result.total
        /*关闭加载*/
        $loadingBar.finish()
      }
  )
}

/*新增文章*/
const showModal = ref(false)

function AddArt() {
  showModal.value = true
  operationType.value = 'add'
}

const formValue = ref(
    {
      title: "",
      desc: "",
      content: "填写内容",
      img: "",
      open: 0,
      sortId: 0,
      id: null
    })
/*查询分类数组*/
const songs = ref([])
/*判断添加编辑变量*/
const operationType = ref('add')

/*编辑和添加文章*/
function addArt() {

  if (operationType.value === 'add') {
    //初始化多余数据
    formValue.value.id = null
    formValue.value.created_at = null
    formValue.value.updated_at = null

    api.addArt(formValue.value).then(
        res => {
          $message.success("添加成功" + res.data.msg)
          // 重新当前页面数据
          page(pagination.page)
          //关闭添加框
          showModal.value = false
        }
    )
  } else {
    api.upArt(formValue.value).then(
        res => {
          $message.success("编辑成功" + res.data.msg)
          // 重新当前页面数据
          page(pagination.page)
          //关闭添加框
          showModal.value = false
        }
    )
  }


}

/*编辑按钮*/
const etArt = (rowData) => {
  operationType.value = "update"
  //开启弹窗
  showModal.value = true
  //赋值给动态表单
  formValue.value = rowData

}

// 处理发布按钮
function handleOpenChange(row) {
  // 更新行数据中的 open 状态
  row.open = row.open === 0 ? 1 : 0;
  //赋值给动态表单
  formValue.value = row

  // 调用服务器 API 更新状态
  operationType.value = "update"
  addArt()

}
</script>

<template>

  <n-card class="h-100% min-w-900 ">
    <!--  筛选 添加框  -->
    <n-card class="flex " style="background-color: rgba(250, 250, 252, 1)">
      <div class="flex items-center w-100%">
        <div class="w-500px ">
          <n-space class="items-center ">
            <span>文章</span>
            <n-input v-model:value="seArt.title" :input-props="{ autocomplete: 'wz' ,id: 'wz'}" placeholder="文章"
                     style="min-width: 250px"
                     type="text"/>
            <n-button type="tertiary" @click="SearchArticle">
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

          <n-button color="#DD5D68CC" @click="DeleteArticle">
            删除
            <template #icon>
              <n-icon
                  :class="'i-ep-semi-select'"
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
        :pagination=pagination
        :row-key="rowKey"
        class="mt-20"
        min-height="690"
        remote
        size="large"
        @update:page="page"
        @update:checked-row-keys="handleCheck"
    />
  </n-card>

  <!--动态弹出框-->
  <n-modal v-model:show="showModal">
    <n-card
        :bordered="false"
        :title="operationType === 'add' ? '添加文章' : '修改文章'"
        aria-modal="true"
        class="w-50% min-w-900"
        role="dialog"
        size="huge"
    >
      <n-form :model="formValue" label-placement="left" label-width="auto">
        <n-form-item label="文章名">
          <n-input v-model:value="formValue.title" placeholder="文章名"/>
        </n-form-item>

        <n-form-item label="描述">
          <n-input v-model:value="formValue.desc" :autosize="{
          minRows: 1,
          maxRows: 5,
        }" placeholder="描述" type="textarea"/>
        </n-form-item>

        <n-form-item label="分类">
          <n-radio-group v-model:value="formValue.sortId">
            <n-space>
              <n-radio v-for="song in songs" :key="song.id" :value="song.id">
                {{ song.name }}
              </n-radio>
            </n-space>
          </n-radio-group>
        </n-form-item>

        <n-form-item label="发布">
          <n-switch v-model:value="formValue.open"
                    :checked-value=1
                    :rail-style="railStyle"
                    :unchecked-value=0
          />
        </n-form-item>

        <n-form-item label="内容">
          <MdEditor v-model="formValue.content"/>
        </n-form-item>
      </n-form>


      <div class="grid place-items-end ">
        <n-button class="w-100" type="info" @click="addArt">
          {{ operationType === 'add' ? '添加' : '保存更改' }}
        </n-button>
      </div>

    </n-card>
  </n-modal>
</template>

<style scoped>

</style>