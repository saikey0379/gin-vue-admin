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
        <el-form-item label="序列号" prop="sn">
         <el-input v-model="searchInfo.sn" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="主机名" prop="hostname">
         <el-input v-model="searchInfo.hostname" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="IP地址" prop="ip">
         <el-input v-model="searchInfo.ip" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="网段ID" prop="network_id">
            
             <el-input v-model.number="searchInfo.network_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="任务负责人" prop="manager">
         <el-input v-model="searchInfo.manager" placeholder="搜索条件" />

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
        <el-table-column align="left" label="序列号" prop="sn" width="120" />
        <el-table-column sortable align="left" label="主机名" prop="hostname" width="120" />
        <el-table-column sortable align="left" label="IP地址" prop="ip" width="120" />
        <el-table-column sortable align="left" label="网段ID" prop="network_id" width="120" />
        <el-table-column sortable align="left" label="硬件配置" prop="hardware_id" width="120" />
        <el-table-column sortable align="left" label="PXE" prop="pxe_id" width="120" />
        <el-table-column sortable align="left" label="Kickstart" prop="kickstart_id" width="120" />
        <el-table-column sortable align="left" label="设备状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,DeviceStatusOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" label="任务负责人" prop="manager" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateOsInstallQueueFunc(scope.row)">变更</el-button>
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
            <el-form-item label="序列号:"  prop="sn" >
              <el-input v-model="formData.sn" :clearable="false"  placeholder="请输入序列号" />
            </el-form-item>
            <el-form-item label="主机名:"  prop="hostname" >
              <el-input v-model="formData.hostname" :clearable="false"  placeholder="请输入主机名" />
            </el-form-item>
            <el-form-item label="IP地址:"  prop="ip" >
              <el-input v-model="formData.ip" :clearable="false"  placeholder="请输入IP地址" />
            </el-form-item>
            <el-form-item label="网段ID:"  prop="network_id" >
              <el-input v-model.number="formData.network_id" :clearable="false" placeholder="请输入网段ID" />
            </el-form-item>
            <el-form-item label="硬件配置:"  prop="hardware_id" >
              <el-input v-model.number="formData.hardware_id" :clearable="false" placeholder="请输入硬件配置" />
            </el-form-item>
            <el-form-item label="PXE:"  prop="pxe_id" >
              <el-input v-model.number="formData.pxe_id" :clearable="false" placeholder="请输入PXE" />
            </el-form-item>
            <el-form-item label="Kickstart:"  prop="kickstart_id" >
              <el-input v-model.number="formData.kickstart_id" :clearable="false" placeholder="请输入Kickstart" />
            </el-form-item>
            <el-form-item label="设备状态:"  prop="status" >
              <el-select v-model="formData.status" placeholder="请选择设备状态" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in DeviceStatusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="任务负责人:"  prop="manager" >
              <el-input v-model="formData.manager" :clearable="false"  placeholder="请输入任务负责人" />
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
                <el-descriptions-item label="序列号">
                        {{ formData.sn }}
                </el-descriptions-item>
                <el-descriptions-item label="主机名">
                        {{ formData.hostname }}
                </el-descriptions-item>
                <el-descriptions-item label="IP地址">
                        {{ formData.ip }}
                </el-descriptions-item>
                <el-descriptions-item label="网段ID">
                        {{ formData.network_id }}
                </el-descriptions-item>
                <el-descriptions-item label="硬件配置">
                        {{ formData.hardware_id }}
                </el-descriptions-item>
                <el-descriptions-item label="PXE">
                        {{ formData.pxe_id }}
                </el-descriptions-item>
                <el-descriptions-item label="Kickstart">
                        {{ formData.kickstart_id }}
                </el-descriptions-item>
                <el-descriptions-item label="设备状态">
                        {{ filterDict(formData.status,DeviceStatusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="任务负责人">
                        {{ formData.manager }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createOsInstallQueue,
  deleteOsInstallQueue,
  deleteOsInstallQueueByIds,
  updateOsInstallQueue,
  findOsInstallQueue,
  getOsInstallQueueList
} from '@/api/osInstallQueue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'OsInstallQueue'
})

// 自动化生成的字典（可能为空）以及字段
const DeviceStatusOptions = ref([])
const formData = ref({
        sn: '',
        hostname: '',
        ip: '',
        network_id: 0,
        hardware_id: 0,
        pxe_id: 0,
        kickstart_id: 0,
        status: undefined,
        manager: '',
        })


// 验证规则
const rule = reactive({
               sn : [{
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
               hostname : [{
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
               ip : [{
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
               network_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               hardware_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               pxe_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               kickstart_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               manager : [{
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
  const table = await getOsInstallQueueList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    DeviceStatusOptions.value = await getDictFunc('DeviceStatus')
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
            deleteOsInstallQueueFunc(row)
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
      const res = await deleteOsInstallQueueByIds({ ids })
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
const updateOsInstallQueueFunc = async(row) => {
    const res = await findOsInstallQueue({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reosInstallQueue
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOsInstallQueueFunc = async (row) => {
    const res = await deleteOsInstallQueue({ ID: row.ID })
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
  const res = await findOsInstallQueue({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reosInstallQueue
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          sn: '',
          hostname: '',
          ip: '',
          network_id: 0,
          hardware_id: 0,
          pxe_id: 0,
          kickstart_id: 0,
          status: undefined,
          manager: '',
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
        sn: '',
        hostname: '',
        ip: '',
        network_id: 0,
        hardware_id: 0,
        pxe_id: 0,
        kickstart_id: 0,
        status: undefined,
        manager: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createOsInstallQueue(formData.value)
                  break
                case 'update':
                  res = await updateOsInstallQueue(formData.value)
                  break
                default:
                  res = await createOsInstallQueue(formData.value)
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
