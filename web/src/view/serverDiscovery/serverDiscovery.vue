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
        <el-form-item label="SN" prop="sn">
         <el-input v-model="searchInfo.sn" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="Hostname" prop="hostname">
         <el-input v-model="searchInfo.hostname" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="IP" prop="ip">
         <el-input v-model="searchInfo.ip" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="操作系统版本" prop="versionOs">
         <el-input v-model="searchInfo.versionOs" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="内核版本" prop="versionKernel">
         <el-input v-model="searchInfo.versionKernel" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="厂商" prop="manufacturer">
         <el-input v-model="searchInfo.manufacturer" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="型号" prop="modelName">
         <el-input v-model="searchInfo.modelName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="CPU" prop="cpuSum">
            
             <el-input v-model.number="searchInfo.cpuSum" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="内存" prop="memorySum">
            
             <el-input v-model.number="searchInfo.memorySum" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="网卡" prop="nicInfo">
         <el-input v-model="searchInfo.nicInfo" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="Gpu" prop="gpu">
         <el-input v-model="searchInfo.gpu" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="主板" prop="motherboard">
         <el-input v-model="searchInfo.motherboard" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="Server类型" prop="serverType">
            <el-select v-model="searchInfo.serverType" clearable placeholder="请选择" @clear="()=>{searchInfo.serverType=undefined}">
              <el-option v-for="(item,key) in ServerTypeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="Agent版本" prop="versionAgent">
         <el-input v-model="searchInfo.versionAgent" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="onEntry">录入</el-button>
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
        <el-table-column align="left" label="日期" width="160">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="SN" prop="sn" width="140" />
        <el-table-column sortable align="left" label="Hostname" prop="hostname" width="140" />
        <el-table-column sortable align="left" label="IP" prop="ip" width="140" />
        <el-table-column sortable align="left" label="操作系统版本" prop="versionOs" width="160" />
        <el-table-column sortable align="left" label="内核版本" prop="versionKernel" width="120" />
        <el-table-column sortable align="left" label="厂商" prop="manufacturer" width="120" />
        <el-table-column align="left" label="型号" prop="modelName" width="120" />
        <el-table-column sortable align="left" label="CPU" prop="cpuSum" width="80" />
        <el-table-column sortable align="left" label="内存" prop="memorySum" width="80" />
        <el-table-column sortable align="left" label="类型" prop="serverType" width="80">
          <template #default="scope">
            {{ filterDict(scope.row.serverType,ServerTypeOptions) }}
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="Agent" prop="versionAgent" width="90" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateServerDiscoveryFunc(scope.row)">变更</el-button>
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
            <el-form-item label="SN:"  prop="sn" >
              <el-input v-model="formData.sn" :clearable="false"  placeholder="请输入SN" />
            </el-form-item>
            <el-form-item label="Hostname:"  prop="hostname" >
              <el-input v-model="formData.hostname" :clearable="false"  placeholder="请输入Hostname" />
            </el-form-item>
            <el-form-item label="IP:"  prop="ip" >
              <el-input v-model="formData.ip" :clearable="false"  placeholder="请输入IP" />
            </el-form-item>
            <el-form-item label="操作系统版本:"  prop="versionOs" >
              <el-input v-model="formData.versionOs" :clearable="false"  placeholder="请输入操作系统版本" />
            </el-form-item>
            <el-form-item label="内核版本:"  prop="versionKernel" >
              <el-input v-model="formData.versionKernel" :clearable="false"  placeholder="请输入内核版本" />
            </el-form-item>
            <el-form-item label="厂商:"  prop="manufacturer" >
              <el-input v-model="formData.manufacturer" :clearable="false"  placeholder="请输入厂商" />
            </el-form-item>
            <el-form-item label="型号:"  prop="modelName" >
              <el-input v-model="formData.modelName" :clearable="false"  placeholder="请输入型号" />
            </el-form-item>
            <el-form-item label="CPU:"  prop="cpuSum" >
              <el-input v-model.number="formData.cpuSum" :clearable="false" placeholder="请输入CPU" />
            </el-form-item>
            <el-form-item label="CPU信息:"  prop="cpu" >
              <el-input v-model="formData.cpu" :clearable="false"  placeholder="请输入CPU信息" />
            </el-form-item>
            <el-form-item label="内存:"  prop="memorySum" >
              <el-input v-model.number="formData.memorySum" :clearable="false" placeholder="请输入内存" />
            </el-form-item>
            <el-form-item label="内存信息:"  prop="memory" >
              <el-input v-model="formData.memory" :clearable="false"  placeholder="请输入内存信息" />
            </el-form-item>
            <el-form-item label="网卡:"  prop="nicInfo" >
              <el-input v-model="formData.nicInfo" :clearable="false"  placeholder="请输入网卡" />
            </el-form-item>
            <el-form-item label="NIC:"  prop="nicDevice" >
              <el-input v-model="formData.nicDevice" :clearable="false"  placeholder="请输入NIC" />
            </el-form-item>
            <el-form-item label="Gpu:"  prop="gpu" >
              <el-input v-model="formData.gpu" :clearable="false"  placeholder="请输入Gpu" />
            </el-form-item>
            <el-form-item label="主板:"  prop="motherboard" >
              <el-input v-model="formData.motherboard" :clearable="false"  placeholder="请输入主板" />
            </el-form-item>
            <el-form-item label="Raid卡:"  prop="raid" >
              <el-input v-model="formData.raid" :clearable="false"  placeholder="请输入Raid卡" />
            </el-form-item>
            <el-form-item label="硬盘信息:"  prop="disk" >
              <el-input v-model="formData.disk" :clearable="false"  placeholder="请输入硬盘信息" />
            </el-form-item>
            <el-form-item label="Server类型:"  prop="serverType" >
              <el-select v-model="formData.serverType" placeholder="请选择Server类型" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in ServerTypeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="Agent版本:"  prop="versionAgent" >
              <el-input v-model="formData.versionAgent" :clearable="false"  placeholder="请输入Agent版本" />
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
                <el-descriptions-item label="SN">
                        {{ formData.sn }}
                </el-descriptions-item>
                <el-descriptions-item label="Hostname">
                        {{ formData.hostname }}
                </el-descriptions-item>
                <el-descriptions-item label="IP">
                        {{ formData.ip }}
                </el-descriptions-item>
                <el-descriptions-item label="操作系统版本">
                        {{ formData.versionOs }}
                </el-descriptions-item>
                <el-descriptions-item label="内核版本">
                        {{ formData.versionKernel }}
                </el-descriptions-item>
                <el-descriptions-item label="厂商">
                        {{ formData.manufacturer }}
                </el-descriptions-item>
                <el-descriptions-item label="型号">
                        {{ formData.modelName }}
                </el-descriptions-item>
                <el-descriptions-item label="CPU">
                        {{ formData.cpuSum }}
                </el-descriptions-item>
                <el-descriptions-item label="CPU信息">
                        {{ formData.cpu }}
                </el-descriptions-item>
                <el-descriptions-item label="内存">
                        {{ formData.memorySum }}
                </el-descriptions-item>
                <el-descriptions-item label="内存信息">
                        {{ formData.memory }}
                </el-descriptions-item>
                <el-descriptions-item label="网卡">
                        {{ formData.nicInfo }}
                </el-descriptions-item>
                <el-descriptions-item label="NIC">
                        {{ formData.nicDevice }}
                </el-descriptions-item>
                <el-descriptions-item label="Gpu">
                        {{ formData.gpu }}
                </el-descriptions-item>
                <el-descriptions-item label="主板">
                        {{ formData.motherboard }}
                </el-descriptions-item>
                <el-descriptions-item label="Raid卡">
                        {{ formData.raid }}
                </el-descriptions-item>
                <el-descriptions-item label="硬盘信息">
                        {{ formData.disk }}
                </el-descriptions-item>
                <el-descriptions-item label="Server类型">
                        {{ filterDict(formData.serverType,ServerTypeOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="Agent版本">
                        {{ formData.versionAgent }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createServerDiscovery,
  deleteServerDiscovery,
  deleteServerDiscoveryByIds,
  updateServerDiscovery,
  findServerDiscovery,
  getServerDiscoveryList,
  entryServerDiscoveryByIds
} from '@/api/serverDiscovery'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'ServerDiscovery'
})

// 自动化生成的字典（可能为空）以及字段
const ServerTypeOptions = ref([])
const formData = ref({
        sn: '',
        hostname: '',
        ip: '',
        versionOs: '',
        versionKernel: '',
        manufacturer: '',
        modelName: '',
        cpuSum: 0,
        cpu: '',
        memorySum: 0,
        memory: '',
        nicInfo: '',
        nicDevice: '',
        gpu: '',
        motherboard: '',
        raid: '',
        disk: '',
        serverType: undefined,
        versionAgent: '',
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
               versionOs : [{
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
               versionKernel : [{
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
               manufacturer : [{
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
               modelName : [{
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
               cpuSum : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               cpu : [{
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
               memorySum : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               memory : [{
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
               nicInfo : [{
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
               nicDevice : [{
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
               gpu : [{
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
               motherboard : [{
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
               raid : [{
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
               disk : [{
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
               serverType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               versionAgent : [{
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
  const table = await getServerDiscoveryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    ServerTypeOptions.value = await getDictFunc('ServerType')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 多选录入
const onEntry = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要录入的数据'
    })
    return
  }
  multipleSelection.value &&
  multipleSelection.value.map(item => {
    ids.push(item.ID)
  })
  const res = await entryServerDiscoveryByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '录入成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteServerDiscoveryFunc(row)
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
      const res = await deleteServerDiscoveryByIds({ ids })
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
const updateServerDiscoveryFunc = async(row) => {
    const res = await findServerDiscovery({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reserverDiscovery
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteServerDiscoveryFunc = async (row) => {
    const res = await deleteServerDiscovery({ ID: row.ID })
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
  const res = await findServerDiscovery({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reserverDiscovery
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
          versionOs: '',
          versionKernel: '',
          manufacturer: '',
          modelName: '',
          cpuSum: 0,
          cpu: '',
          memorySum: 0,
          memory: '',
          nicInfo: '',
          nicDevice: '',
          gpu: '',
          motherboard: '',
          raid: '',
          disk: '',
          serverType: undefined,
          versionAgent: '',
          }
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        sn: '',
        hostname: '',
        ip: '',
        versionOs: '',
        versionKernel: '',
        manufacturer: '',
        modelName: '',
        cpuSum: 0,
        cpu: '',
        memorySum: 0,
        memory: '',
        nicInfo: '',
        nicDevice: '',
        gpu: '',
        motherboard: '',
        raid: '',
        disk: '',
        serverType: undefined,
        versionAgent: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createServerDiscovery(formData.value)
                  break
                case 'update':
                  res = await updateServerDiscovery(formData.value)
                  break
                default:
                  res = await createServerDiscovery(formData.value)
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
