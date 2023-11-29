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
        <el-form-item label="名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="描述" prop="description">
         <el-input v-model="searchInfo.description" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="负责人" prop="manager">
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
        <el-table-column sortable align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="描述" prop="description" width="120" />
        <el-table-column sortable align="left" label="负责人" prop="manager" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateFileBinaryFunc(scope.row)">变更</el-button>
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
            <el-form-item label="名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="false"  placeholder="请输入名称" />
            </el-form-item>
            <el-form-item label="描述:"  prop="description" >
              <el-input v-model="formData.description" :clearable="false"  placeholder="请输入描述" />
            </el-form-item>
            <el-form-item label="文件:" :rules="[{ required: true }]" >
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
            </el-form-item>
            <el-form-item label="进度"  prop="percent" >
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
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button  v-if="percentage === 100" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions :column="1" border>
                <el-descriptions-item label="名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="描述">
                        {{ formData.description }}
                </el-descriptions-item>
                <el-descriptions-item label="MD5">
                        {{ formData.md5 }}
                </el-descriptions-item>
                <el-descriptions-item label="链接">
                        {{ formData.url }}
                </el-descriptions-item>
                <el-descriptions-item label="负责人">
                        {{ formData.manager }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createFileBinary,
  deleteFileBinary,
  deleteFileBinaryByIds,
  updateFileBinary,
  findFileBinary,
  getFileBinaryList
} from '@/api/fileBinary'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch } from 'vue'

defineOptions({
    name: 'FileBinary',
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        name: '',
        description: '',
        md5: '',
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
                   message: '若不填写，默认为上传文件名',
                   trigger: ['input', 'blur'],
              }
              ],
               description : [{
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
  const table = await getFileBinaryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteFileBinaryFunc(row)
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
      const res = await deleteFileBinaryByIds({ ids })
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
const updateFileBinaryFunc = async(row) => {
    const res = await findFileBinary({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.refileBinary
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteFileBinaryFunc = async (row) => {
    const res = await deleteFileBinary({ ID: row.ID })
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
  const res = await findFileBinary({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.refileBinary
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          name: '',
          description: '',
          md5: '',
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
        description: '',
        md5: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createFileBinary(formData.value)
                  break
                case 'update':
                  res = await updateFileBinary(formData.value)
                  break
                default:
                  res = await createFileBinary(formData.value)
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

import SparkMD5 from 'spark-md5'
import {
  findFile,
  breakpointContinueFinish,
  removeChunk,
  breakpointContinue,
} from '@/api/breakpoint'

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

  if (formData.value.name !== '') {
    fileName.value = formData.value.name
  } else {
    fileName.value = file.value.name
    formData.value.name = file.value.name
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
    const res = await findFile(params)
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
      upLoadFileSlice(item)
    }
  })
}

watch(() => waitNum.value, () => { percentage.value = Math.floor(((formDataList.value.length - waitNum.value) / formDataList.value.length) * 100) })

const upLoadFileSlice = async(item) => {
  // 切片上传
  const fileRe = await breakpointContinue(item.formData)
  if (fileRe.code !== 0) {
    return
  }
  waitNum.value-- // 百分数增加
  if (waitNum.value === 0) {
    // 切片传完以后 合成文件
    const params = {
      fileName: fileName.value,
      fileMd5: fileMd5.value
    }
    const res = await breakpointContinueFinish(params)
    if (res.code === 0) {
      // 合成文件过后 删除缓存切片
      const params = {
        fileName: fileName.value,
        fileMd5: fileMd5.value,
        filePath: res.data.filePath,
      }
      ElMessage.success('上传成功')
      await removeChunk(params)
    }
  }
}

const FileInput = ref(null)
const inputChange = () => {
  FileInput.value.dispatchEvent(new MouseEvent('click'))
}
</script>

<style lang='scss' scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#fromCont{
  display: inline-block;
}
.fileUpload{
  padding: 3px 10px;
  font-size: 12px;
  height: 20px;
  line-height: 20px;
  position: relative;
  cursor: pointer;
  color: #000;
  border: 1px solid #c1c1c1;
  border-radius: 4px;
  overflow: hidden;
  display: inline-block;
  input{
    position: absolute;
    font-size: 100px;
    right: 0;
    top: 0;
    opacity: 0;
    cursor: pointer;
  }
}
.uploadBtn{
  position: relative;
  top: -5px;
  margin-left: 15px;
  width: 45px; /* 设置按钮的宽度 */
  height: 25px; /* 设置按钮的高度 */
}
.list-item {
  list-style: none;
  display: block;
  margin-right: 10px;
  color: #606266;
  line-height: 25px;
  width: 40%;
  .percentage{
    float: right;
  }
}

</style>
