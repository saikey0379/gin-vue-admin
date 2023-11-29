<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item label="Name" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="负责人" prop="manager">
         <el-input v-model="searchInfo.manager" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="节点" prop="nodes">
         <el-input v-model="searchInfo.nodes" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="文件类型" prop="fileType">
            <el-select v-model="searchInfo.fileType" clearable placeholder="请选择" @clear="()=>{searchInfo.fileType=undefined}">
              <el-option v-for="(item,key) in FileTypeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="文件ID" prop="fileId">
            
             <el-input v-model.number="searchInfo.fileId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="目标路径" prop="destPath">
         <el-input v-model="searchInfo.destPath" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="Name" prop="name" width="120" />
        <el-table-column sortable align="left" label="负责人" prop="manager" width="120" />
        <el-table-column align="left" label="描述" prop="describe" width="120" />
        <el-table-column align="left" label="节点" prop="nodes" width="120" />
        <el-table-column sortable align="left" label="文件类型" prop="fileType" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.fileType,FileTypeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="文件ID" prop="fileId" width="120" />
        <el-table-column align="left" label="文件权限" prop="fileMod" width="120" />
        <el-table-column align="left" label="执行参数" prop="parameter" width="120" />
        <el-table-column align="left" label="目标路径" prop="destPath" width="120" />
        <el-table-column align="left" label="Minute" prop="cronMinute" width="120" />
        <el-table-column align="left" label="Hour" prop="cronHour" width="120" />
        <el-table-column align="left" label="DayOfMonth" prop="cronDayOfMonth" width="120" />
        <el-table-column align="left" label="Month" prop="cronMonth" width="120" />
        <el-table-column align="left" label="DayOfWeek" prop="cronDayOfWeek" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateTaskCrontabFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="Name:"  prop="name" >
              <el-input v-model="formData.name" :clearable="false"  placeholder="请输入Name" />
            </el-form-item>
            <el-form-item label="负责人:"  prop="manager" >
              <el-input v-model="formData.manager" :clearable="false"  placeholder="请输入负责人" />
            </el-form-item>
            <el-form-item label="描述:"  prop="describe" >
              <el-input v-model="formData.describe" :clearable="false"  placeholder="请输入描述" />
            </el-form-item>
            <el-form-item label="节点:"  prop="nodes" >
              <el-input v-model="formData.nodes" :clearable="false"  placeholder="请输入节点" />
            </el-form-item>
            <el-form-item label="文件类型:"  prop="fileType" >
              <el-select v-model="formData.fileType" placeholder="请选择文件类型" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in FileTypeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="文件ID:"  prop="fileId" >
              <el-input v-model.number="formData.fileId" :clearable="false" placeholder="请输入文件ID" />
            </el-form-item>
            <el-form-item label="文件权限:"  prop="fileMod" >
              <el-input v-model="formData.fileMod" :clearable="false"  placeholder="请输入文件权限" />
            </el-form-item>
            <el-form-item label="执行参数:"  prop="parameter" >
              <el-input v-model="formData.parameter" :clearable="false"  placeholder="请输入执行参数" />
            </el-form-item>
            <el-form-item label="目标路径:"  prop="destPath" >
              <el-input v-model="formData.destPath" :clearable="false"  placeholder="请输入目标路径" />
            </el-form-item>
            <el-form-item label="Minute:"  prop="cronMinute" >
              <el-input v-model="formData.cronMinute" :clearable="false"  placeholder="请输入Minute" />
            </el-form-item>
            <el-form-item label="Hour:"  prop="cronHour" >
              <el-input v-model="formData.cronHour" :clearable="false"  placeholder="请输入Hour" />
            </el-form-item>
            <el-form-item label="DayOfMonth:"  prop="cronDayOfMonth" >
              <el-input v-model="formData.cronDayOfMonth" :clearable="true"  placeholder="请输入DayOfMonth" />
            </el-form-item>
            <el-form-item label="Month:"  prop="cronMonth" >
              <el-input v-model="formData.cronMonth" :clearable="true"  placeholder="请输入Month" />
            </el-form-item>
            <el-form-item label="DayOfWeek:"  prop="cronDayOfWeek" >
              <el-input v-model="formData.cronDayOfWeek" :clearable="false"  placeholder="请输入DayOfWeek" />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="Name">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="负责人">
                        {{ formData.manager }}
                </el-descriptions-item>
                <el-descriptions-item label="描述">
                        {{ formData.describe }}
                </el-descriptions-item>
                <el-descriptions-item label="节点">
                        {{ formData.nodes }}
                </el-descriptions-item>
                <el-descriptions-item label="文件类型">
                        {{ filterDict(formData.fileType,FileTypeOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="文件ID">
                        {{ formData.fileId }}
                </el-descriptions-item>
                <el-descriptions-item label="文件权限">
                        {{ formData.fileMod }}
                </el-descriptions-item>
                <el-descriptions-item label="执行参数">
                        {{ formData.parameter }}
                </el-descriptions-item>
                <el-descriptions-item label="目标路径">
                        {{ formData.destPath }}
                </el-descriptions-item>
                <el-descriptions-item label="Minute">
                        {{ formData.cronMinute }}
                </el-descriptions-item>
                <el-descriptions-item label="Hour">
                        {{ formData.cronHour }}
                </el-descriptions-item>
                <el-descriptions-item label="DayOfMonth">
                        {{ formData.cronDayOfMonth }}
                </el-descriptions-item>
                <el-descriptions-item label="Month">
                        {{ formData.cronMonth }}
                </el-descriptions-item>
                <el-descriptions-item label="DayOfWeek">
                        {{ formData.cronDayOfWeek }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createTaskCrontab,
  deleteTaskCrontab,
  deleteTaskCrontabByIds,
  updateTaskCrontab,
  findTaskCrontab,
  getTaskCrontabList
} from '@/api/taskCrontab'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'TaskCrontab'
})

// 自动化生成的字典（可能为空）以及字段
const FileTypeOptions = ref([])
const formData = ref({
        name: '',
        manager: '',
        describe: '',
        nodes: '',
        fileType: undefined,
        fileId: 0,
        fileMod: '',
        parameter: '',
        destPath: '',
        cronMinute: '',
        cronHour: '',
        cronDayOfMonth: '',
        cronMonth: '',
        cronDayOfWeek: '',
        })


// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               describe : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               nodes : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               fileType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               fileId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               cronMinute : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               cronHour : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               cronDayOfMonth : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               cronMonth : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               cronDayOfWeek : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getTaskCrontabList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    FileTypeOptions.value = await getDictFunc('FileType')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteTaskCrontabFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteTaskCrontabByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTaskCrontabFunc = async(row) => {
    const res = await findTaskCrontab({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.retaskCrontab
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTaskCrontabFunc = async (row) => {
    const res = await deleteTaskCrontab({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findTaskCrontab({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.retaskCrontab
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          name: '',
          manager: '',
          describe: '',
          nodes: '',
          fileType: undefined,
          fileId: 0,
          fileMod: '',
          parameter: '',
          destPath: '',
          cronMinute: '',
          cronHour: '',
          cronDayOfMonth: '',
          cronMonth: '',
          cronDayOfWeek: '',
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        name: '',
        manager: '',
        describe: '',
        nodes: '',
        fileType: undefined,
        fileId: 0,
        fileMod: '',
        parameter: '',
        destPath: '',
        cronMinute: '',
        cronHour: '',
        cronDayOfMonth: '',
        cronMonth: '',
        cronDayOfWeek: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createTaskCrontab(formData.value)
                  break
                case 'update':
                  res = await updateTaskCrontab(formData.value)
                  break
                default:
                  res = await createTaskCrontab(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
