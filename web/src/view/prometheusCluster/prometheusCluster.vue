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
        >
        <el-table-column type="selection" width="55" />
        <el-table-column sortable align="left" label="Name" prop="name" width="180" />
        <el-table-column align="left" label="用户名" prop="sshUser" width="150" />
        <el-table-column align="left" label="节点" prop="nodes" width="200" />
        <el-table-column sortable align="left" label="证书" prop="sshRsa" width="200" />
        <el-table-column sortable align="left" label="负责人" prop="manager" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link class="table-button" @click="getPromYamlListByClusterId(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看配置
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updatePrometheusClusterFunc(scope.row)">变更</el-button>
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
            <el-form-item label="描述:"  prop="describe" >
              <el-input v-model="formData.describe" :clearable="false"  placeholder="请输入描述" />
            </el-form-item>
            <el-form-item label="用户名:"  prop="sshUser" >
              <el-input v-model="formData.sshUser" :clearable="false"  placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item label="RSA:"  prop="sshRsa" >
              <div style="display: flex; align-items: center;">
              <el-input v-model="formData.sshRsa" :clearable="false"  placeholder="请输入RSA文件名" />
            <form id="fromCont" method="post" >
              <div class="fileUpload" @click="inputChange" >
                选择文件
                <input
                    v-show="false"
                    id="file"
                    ref="FileInput"
                    multiple="multiple"
                    type="file"
                    @change="choseFile"
                >
              </div>
            </form>
            <el-button
                :disabled="limitFileSize"
                type="primary"
                class="uploadBtn"
                @click="getFile"
            >上传</el-button>
              <transition name="list" tag="p" >
                <div v-if="file" class="list-item" >
                  <span>{{ file.name }}</span>
                  <span class="percentage">{{ percentage }}%</span>
                  <el-progress
                      :show-text="false"
                      :text-inside="false"
                      :stroke-width="6"
                      :percentage="percentage"
                  />
                </div>
              </transition>
              </div>
            </el-form-item>
            <el-form-item label="节点:"  prop="nodes" >
              <el-input v-model="formData.nodes" :clearable="false"  placeholder="请输入节点" />
            </el-form-item>
            <el-form-item label="配置路径:"  prop="pathConf" >
              <el-input v-model="formData.pathConf" :clearable="false"  placeholder="请输入配置路径" />
            </el-form-item>
            <el-form-item label="加载命令:"  prop="execLoad" >
              <el-input v-model="formData.execLoad" :clearable="false"  placeholder="请输入加载命令" />
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
                <el-descriptions-item label="用户名">
                        {{ formData.sshUser }}
                </el-descriptions-item>
                <el-descriptions-item label="RSA">
                        {{ formData.sshRsa }}
                </el-descriptions-item>
                <el-descriptions-item label="节点">
                        {{ formData.nodes }}
                </el-descriptions-item>
                <el-descriptions-item label="配置路径">
                        {{ formData.pathConf }}
                </el-descriptions-item>
                <el-descriptions-item label="加载命令">
                        {{ formData.execLoad }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
    <el-dialog v-model="showPromYaml" width="80%" :before-close="closeShowRuleYaml" title="配置详情"  destroy-on-close>
      <template #header>
        <div style="display: flex; justify-content: space-between;">
          <div>配置详情</div>
          <div>
            <el-button type="info" @click="handlePromYamlModeList(promYaml.ID)">列表模式</el-button>
            <el-button type="info" @click="handlePromYamlModeDetail(promYaml.ID)">全局模式</el-button>
            <el-button type="primary" @click="handlePromYamlModeDetailCurrent(promYaml.ID)">实时配置</el-button>
            <el-button type="primary" @click="handlePromYamlModeDiff(promYaml.ID)">DIFF</el-button>
            <el-button type="danger" @click="updatePrometheusRulesConfirm(promYaml.ID)">更新</el-button>
          </div>
        </div>
      </template>
        <div v-if="!showPromYamlDiff" class="rule-group-title">规则组数量: {{ promYaml.totalGroups }}</div>
        <div v-if="showPromYamlList" v-for="(rule, index) in promYaml.groups" :key="index" style="margin-bottom: 20px;">
          <div class="rule-group">
            <div class="rule-group-list">
              <div class="rule-group-name" style="display: flex; align-items: center;">
                <p style="margin-right: 10px;">RuleGroup:</p>
                <p>{{ rule.name }}</p>
                <el-tooltip content="复制">
                  <el-icon type="link"  @click="copyToClipboard(rule.name)"><CopyDocument /></el-icon>
                </el-tooltip>
              </div>
              <div class="rule-count" style="display: flex; ">
                <p style="margin-right: 10px;  solid #000;">RuleTotal:</p>
                <p>{{ rule.totalRules }}</p>
              </div>
            </div>
            <div style="position: relative;">
              <pre class="rule-yaml">{{ rule.content }}</pre>
              <el-tooltip content="复制">
                <el-icon type="link" :style="{ color: 'white', position: 'absolute', top: '0', right: '0' }" @click="copyToClipboard(rule.content)"><CopyDocument /></el-icon>
              </el-tooltip>
            </div>
          </div>
        </div>
        <div v-if="showPromYamlDetail" style="position: relative;">
          <div class="rule-group">
            <div class="rule-group-list">
              <div class="rule-count" style="display: flex; ">
                <p style="margin-right: 10px;  solid #000;">RuleTotal:</p>
                <p>{{ promYaml.totalRules }}</p>
              </div>
            </div>
            <div style="position: relative;">
              <pre class="rule-yaml">{{ promYaml.content }}</pre>
              <el-tooltip content="复制">
                <el-icon type="link" :style="{ color: 'white', position: 'absolute', top: '0', right: '0' }" @click="copyToClipboard(promYaml.content)"><CopyDocument /></el-icon>
              </el-tooltip>
            </div>
          </div>
        </div>
        <div v-if="showPromYamlDiff" style="position: relative;">
          <div class="rule-group-title">目标规则组数量: {{ promYaml.totalGroups }}</div>
          <div class="rule-group-title">目标规则数量: {{ promYaml.totalRules }}</div>
          <div class="rule-group-title">当前规则组数量: {{ promYaml.totalCurrentGroups }}</div>
          <div class="rule-group-title">当前规则数量: {{ promYaml.totalCurrentRules }}</div>
          <div class="rule-group">
            <div style="position: relative;">
              <pre class="rule-yaml">{{ promYaml.content }}</pre>
              <el-tooltip content="复制">
                <el-icon type="link" :style="{ color: 'white', position: 'absolute', top: '0', right: '0' }" @click="copyToClipboard(promYaml.content)"><CopyDocument /></el-icon>
              </el-tooltip>
            </div>
          </div>
        </div>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createPrometheusCluster,
  deletePrometheusCluster,
  deletePrometheusClusterByIds,
  updatePrometheusCluster,
  findPrometheusCluster,
  getPrometheusClusterList,
} from '@/api/prometheusCluster'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile, copyToClipboard } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch } from 'vue'
import { findPromYamlListByClusterId,
  findPromYamlDetailByClusterId,
  findPromYamlDetailCurrentByClusterId,
  diffPromYamlByClusterId,
  updatePrometheusRulesByClusterId 
} from '@/api/prometheusCluster'

defineOptions({
    name: 'PrometheusCluster'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        name: '',
        manager: '',
        describe: '',
        sshUser: '',
        sshRsa: '',
        nodes: '',
        pathConf: '',
        execLoad: '',
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
               sshUser : [{
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
               sshRsa : [{
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
               pathConf : [{
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
               execLoad : [{
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
  const table = await getPrometheusClusterList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deletePrometheusClusterFunc(row)
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
      const res = await deletePrometheusClusterByIds({ ids })
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
const updatePrometheusClusterFunc = async(row) => {
    const res = await findPrometheusCluster({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reprometheusCluster
        dialogFormVisible.value = true
    }
}


// 删除行
const deletePrometheusClusterFunc = async (row) => {
    const res = await deletePrometheusCluster({ ID: row.ID })
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
  const res = await findPrometheusCluster({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reprometheusCluster
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
          sshUser: '',
          sshRsa: '',
          nodes: '',
          pathConf: '',
          execLoad: '',
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
        sshUser: '',
        sshRsa: '',
        nodes: '',
        pathConf: '',
        execLoad: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createPrometheusCluster(formData.value)
                  break
                case 'update':
                  res = await updatePrometheusCluster(formData.value)
                  break
                default:
                  res = await createPrometheusCluster(formData.value)
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

// PromYamlList
const promYaml = ref("");

const showPromYaml = ref(false);

const closeShowRuleYaml = () => {
  promYaml.value = ''
  showPromYaml.value = false
}

// PromYamlList
const showPromYamlList = ref(false);

const openShowPromYamlList = () => {
  showPromYaml.value = true
  showPromYamlList.value = true
  showPromYamlDetail.value = false
  showPromYamlDiff.value = false
}

const getPromYamlListByClusterId = async(row) => {
  const res = await findPromYamlListByClusterId({ ID: row.ID })
  promYaml.value = res.data.promYamlCluster
  openShowPromYamlList()
}

const handlePromYamlModeList = async(id) => {
  const res = await findPromYamlListByClusterId({ ID: id })
  promYaml.value = res.data.promYamlCluster
  openShowPromYamlList()
}

// PromYamlDetail
const showPromYamlDetail = ref(false);

const openShowPromYamlDetail = () => {
  showPromYaml.value = true
  showPromYamlList.value = false
  showPromYamlDiff.value = false
  showPromYamlDetail.value = true
}

const handlePromYamlModeDetail = async(id) => {
  const res = await findPromYamlDetailByClusterId({ ID: id })
  promYaml.value = res.data.promYamlCluster
  openShowPromYamlDetail()
}

const handlePromYamlModeDetailCurrent = async(id) => {
  const res = await findPromYamlDetailCurrentByClusterId({ ID: id })
  promYaml.value = res.data.promYamlCluster
  openShowPromYamlDetail()
}

// DiffPromYaml
const showPromYamlDiff = ref(false);

const openShowPromYamlDiff = () => {
  showPromYaml.value = true
  showPromYamlList.value = false
  showPromYamlDetail.value = false
  showPromYamlDiff.value = true
}

const handlePromYamlModeDiff = async(id) => {
  const res = await diffPromYamlByClusterId({ ID: id })
  promYaml.value = res.data.promYamlCluster
  openShowPromYamlDiff()
}

const updatePrometheusRulesConfirm = async(id) => {
  showPromYaml.value = false,
  ElMessageBox.confirm('确定要更新吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(async() => {
      const res = await updatePrometheusRulesByClusterId({ ID: id })
        if (res.code === 0) {
            ElMessage({
                type: 'success',
                message: '更新成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
              page.value--
            }
        }
    })
}

import SparkMD5 from 'spark-md5'
import {
  resumeTransFindFile,
  resumeTransChunkContinue,
  resumeTransChunkFinish,
  resumeTransChunkRemove,
} from '@/api/fileHandle'

const file = ref(null)
const fileName = ref('')
const fileMd5 = ref('')
const formDataList = ref([])
const waitUpLoad = ref([])
const waitNum = ref(NaN)
const limitFileSize = ref(false)
const percentage = ref(0)
const percentageFlage = ref(true)

// 选中文件的函数
const choseFile = async(e) => {
  const fileInput = e.target.files[0] // 获取当前文件
  const maxSize = 4096 * 1024 * 1024
  file.value = fileInput // file 丢全局方便后面用 可以改进为func传参形式

  percentage.value = 0
  if (file.value.size > maxSize) {
    limitFileSize.value = true
    ElMessage('请上传小于4G文件')
  }
}

const getFile = () => {
  const fileR = new FileReader() // 创建一个reader用来读取文件流

  if (formData.value.sshRsa !== '') {
    fileName.value = formData.value.sshRsa
  } else {
    fileName.value = file.value.name
    formData.value.sshRsa = file.value.name
  }

  // 确定按钮
  if (file.value === null) {
    ElMessage('请先上传文件')
    return
  }
  if (percentage.value === 100) {
    percentageFlage.value = false
  }

  limitFileSize.value = false
  fileR.readAsArrayBuffer(file.value) // 把文件读成ArrayBuffer  主要为了保持跟后端的流一致
  fileR.onload = async e => {
    // 读成arrayBuffer的回调 e 为方法自带参数 相当于 dom的e 流存在e.target.result 中
    const blob = e.target.result
    const spark = new SparkMD5.ArrayBuffer() // 创建md5制造工具 （md5用于检测文件一致性 这里不懂就打电话问我）
    spark.append(blob) // 文件流丢进工具
    fileMd5.value = spark.end() // 工具结束 产生一个a 总文件的md5
    formData.value.md5 = fileMd5.value
    const FileSliceCap = 4 * 1024 * 1024 // 分片字节数
    let start = 0 // 定义分片开始切的地方
    let end = 0 // 每片结束切的地方a
    let i = 0 // 第几片
    formDataList.value = [] // 分片存储的一个池子 丢全局
    while (end < file.value.size) {
      // 当结尾数字大于文件总size的时候 结束切片
      start = i * FileSliceCap // 计算每片开始位置
      end = (i + 1) * FileSliceCap // 计算每片结束位置
      var fileSlice = file.value.slice(start, end) // 开始切  file.slice 为 h5方法 对文件切片 参数为 起止字节数
      const formData = new window.FormData() // 创建FormData用于存储传给后端的信息
      formData.append('fileMd5', fileMd5.value) // 存储总文件的Md5 让后端知道自己是谁的切片
      formData.append('file', fileSlice) // 当前的切片
      formData.append('chunkNumber', i) // 当前是第几片
      formData.append('fileName', fileName.value) // 当前文件的文件名 用于后端文件切片的命名  formData.appen 为 formData对象添加参数的方法
      formDataList.value.push({ key: i, formData }) // 把当前切片信息 自己是第几片 存入我们方才准备好的池子
      i++
    }
    const params = {
      fileName: file.value.name,
      fileMd5: fileMd5.value,
      chunkTotal: formDataList.value.length
    }
    const res = await resumeTransFindFile(params)
    // 全部切完以后 发一个请求给后端 拉当前文件后台存储的切片信息 用于检测有多少上传成功的切片
    const finishList = res.data.file.ExaFileChunk // 上传成功的切片
    const IsFinish = res.data.file.IsFinish // 是否是同文件不同命 （文件md5相同 文件名不同 则默认是同一个文件但是不同文件名 此时后台数据库只需要拷贝一下数据库文件即可 不需要上传文件 即秒传功能）
    if (!IsFinish) {
      // 当是断点续传时候
      waitUpLoad.value = formDataList.value.filter(all => {
        return !(
            finishList &&
            finishList.some(fi => fi.FileChunkNumber === all.key)
        ) // 找出需要上传的切片
      })
      sliceFile() // 上传切片
    } else {
      waitUpLoad.value = [] // 秒传则没有需要上传的切片
      percentage.value = 100
      ElMessage.success('文件已秒传')
    }
    waitNum.value = waitUpLoad.value.length // 记录长度用于百分比展示
    console.log(waitNum.value)
  }
}

const sliceFile = () => {
  waitUpLoad.value &&
  waitUpLoad.value.forEach(item => {
    // 需要上传的切片
    item.formData.append('chunkTotal', formDataList.value.length) // 切片总数携带给后台 总有用的
    const fileR = new FileReader() // 功能同上
    const fileF = item.formData.get('file')
    fileR.readAsArrayBuffer(fileF)
    fileR.onload = e => {
      const spark = new SparkMD5.ArrayBuffer()
      spark.append(e.target.result)
      item.formData.append('chunkMd5', spark.end()) // 获取当前切片md5 后端用于验证切片完整性
      item.formData.append('fileType', "rsa")
      upLoadFileSlice(item)
    }
  })
}

const upLoadFileSlice = async(item) => {
  // 切片上传
  const fileRe = await resumeTransChunkContinue(item.formData)
  if (fileRe.code !== 0) {
    return
  }
  waitNum.value-- // 百分数增加
  if (waitNum.value === 0) {
    // 切片传完以后 合成文件
    const params = {
      fileName: fileName.value,
      fileMd5: fileMd5.value,
      fileType: "rsa",
    }
    const res = await resumeTransChunkFinish(params)
    if (res.code === 0) {
      // 合成文件过后 删除缓存切片
      const params = {
        fileName: fileName.value,
        fileMd5: fileMd5.value,
        filePath: res.data.filePath,
      }
      ElMessage.success('上传成功')
      await resumeTransChunkRemove(params)
    }
  }
}

const FileInput = ref(null)
const inputChange = () => {
  FileInput.value.dispatchEvent(new MouseEvent('click'))
}

</script>

<style>
.rule-yaml {
  padding: 15px;
  background-color: black;
  color: white;
  border-radius: 5px;
  white-space: pre-wrap;
  word-wrap: break-word;
  overflow: auto; 
}
.rule-group-title {
  padding: 8px;
  margin-bottom: 2px;
  background-color: #f5f5f5;
  font-weight: bold;
}

.rule-group {
  padding: 8px;
  background-color: #f5f5f5;
}

.fileUpload{
  white-space: nowrap;
  padding: 3px 12px;
  font-size: 12px;
  height: 24px;
  line-height: 25px;
  border: 1px solid #c1c1c1;
  border-radius: 8px;
}
.uploadBtn{
  font-size: 12px;
  margin-left: 2px;
  width: 45px; /* 设置按钮的宽度 */
  height: 28px; /* 设置按钮的高度 */
  line-height: 24px;
}
</style>
