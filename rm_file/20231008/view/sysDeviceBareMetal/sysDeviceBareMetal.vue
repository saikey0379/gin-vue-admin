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
        <el-form-item label="资产编号" prop="asset_id">
         <el-input v-model="searchInfo.asset_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="主机名" prop="hostname">
         <el-input v-model="searchInfo.hostname" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="IP地址" prop="ip">
         <el-input v-model="searchInfo.ip" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="管理IP" prop="manage_ip">
         <el-input v-model="searchInfo.manage_ip" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="网段ID" prop="network_id">
            
             <el-input v-model.number="searchInfo.network_id" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="管理网段ID" prop="manage_network_id">
            
             <el-input v-model.number="searchInfo.manage_network_id" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="硬件配置模版" prop="hardware_id">
            <el-select v-model="searchInfo.hardware_id" clearable placeholder="请选择" @clear="()=>{searchInfo.hardware_id=undefined}">
              <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="PXE模版" prop="pxe_id">
            <el-select v-model="searchInfo.pxe_id" clearable placeholder="请选择" @clear="()=>{searchInfo.pxe_id=undefined}">
              <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="Kickstart模版" prop="kickstart_id">
            <el-select v-model="searchInfo.kickstart_id" clearable placeholder="请选择" @clear="()=>{searchInfo.kickstart_id=undefined}">
              <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="机柜ID" prop="cabinet_id">
            <el-select v-model="searchInfo.cabinet_id" clearable placeholder="请选择" @clear="()=>{searchInfo.cabinet_id=undefined}">
              <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="机柜U位" prop="cabinet_u">
         <el-input v-model="searchInfo.cabinet_u" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="开发负责人" prop="manager_dev">
         <el-input v-model="searchInfo.manager_dev" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="运维负责人" prop="manager_ops">
         <el-input v-model="searchInfo.manager_ops" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="设备标签ID" prop="label_id">
            <el-select v-model="searchInfo.label_id" clearable placeholder="请选择" @clear="()=>{searchInfo.label_id=undefined}">
              <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="设备描述" prop="device_describe">
         <el-input v-model="searchInfo.device_describe" placeholder="搜索条件" />

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
        <el-table-column sortable align="left" label="资产编号" prop="asset_id" width="120" />
        <el-table-column sortable align="left" label="主机名" prop="hostname" width="120" />
        <el-table-column sortable align="left" label="IP地址" prop="ip" width="120" />
        <el-table-column sortable align="left" label="管理IP" prop="manage_ip" width="120" />
        <el-table-column sortable align="left" label="网段ID" prop="network_id" width="120" />
        <el-table-column sortable align="left" label="管理网段ID" prop="manage_network_id" width="120" />
        <el-table-column sortable align="left" label="硬件配置模版" prop="hardware_id" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.hardware_id,intOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" label="PXE模版" prop="pxe_id" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.pxe_id,intOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" label="Kickstart模版" prop="kickstart_id" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.kickstart_id,intOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" label="机柜ID" prop="cabinet_id" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.cabinet_id,intOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" label="机柜U位" prop="cabinet_u" width="120" />
        <el-table-column sortable align="left" label="设备状态" prop="status" width="120" />
        <el-table-column sortable align="left" label="开发负责人" prop="manager_dev" width="120" />
        <el-table-column sortable align="left" label="运维负责人" prop="manager_ops" width="120" />
        <el-table-column sortable align="left" label="设备标签ID" prop="label_id" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.label_id,intOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="设备描述" prop="device_describe" width="120" />
        <el-table-column align="left" label="操作">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBareMetalFunc(scope.row)">变更</el-button>
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
            <el-form-item label="资产编号:"  prop="asset_id" >
              <el-input v-model="formData.asset_id" :clearable="false"  placeholder="请输入资产编号" />
            </el-form-item>
            <el-form-item label="主机名:"  prop="hostname" >
              <el-input v-model="formData.hostname" :clearable="false"  placeholder="请输入主机名" />
            </el-form-item>
            <el-form-item label="IP地址:"  prop="ip" >
              <el-input v-model="formData.ip" :clearable="false"  placeholder="请输入IP地址" />
            </el-form-item>
            <el-form-item label="管理IP:"  prop="manage_ip" >
              <el-input v-model="formData.manage_ip" :clearable="false"  placeholder="请输入管理IP" />
            </el-form-item>
            <el-form-item label="网段ID:"  prop="network_id" >
              <el-input v-model.number="formData.network_id" :clearable="false" placeholder="请输入网段ID" />
            </el-form-item>
            <el-form-item label="管理网段ID:"  prop="manage_network_id" >
              <el-input v-model.number="formData.manage_network_id" :clearable="false" placeholder="请输入管理网段ID" />
            </el-form-item>
            <el-form-item label="硬件配置模版:"  prop="hardware_id" >
              <el-select v-model="formData.hardware_id" placeholder="请选择硬件配置模版" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="PXE模版:"  prop="pxe_id" >
              <el-select v-model="formData.pxe_id" placeholder="请选择PXE模版" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="Kickstart模版:"  prop="kickstart_id" >
              <el-select v-model="formData.kickstart_id" placeholder="请选择Kickstart模版" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="机柜ID:"  prop="cabinet_id" >
              <el-select v-model="formData.cabinet_id" placeholder="请选择机柜ID" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="机柜U位:"  prop="cabinet_u" >
              <el-input v-model="formData.cabinet_u" :clearable="false"  placeholder="请输入机柜U位" />
            </el-form-item>
            <el-form-item label="设备状态:"  prop="status" >
              <el-input v-model="formData.status" :clearable="false"  placeholder="请输入设备状态" />
            </el-form-item>
            <el-form-item label="开发负责人:"  prop="manager_dev" >
              <el-input v-model="formData.manager_dev" :clearable="false"  placeholder="请输入开发负责人" />
            </el-form-item>
            <el-form-item label="运维负责人:"  prop="manager_ops" >
              <el-input v-model="formData.manager_ops" :clearable="false"  placeholder="请输入运维负责人" />
            </el-form-item>
            <el-form-item label="设备标签ID:"  prop="label_id" >
              <el-select v-model="formData.label_id" placeholder="请选择设备标签ID" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="设备描述:"  prop="device_describe" >
              <el-input v-model="formData.device_describe" :clearable="false"  placeholder="请输入设备描述" />
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
                <el-descriptions-item label="资产编号">
                        {{ formData.asset_id }}
                </el-descriptions-item>
                <el-descriptions-item label="主机名">
                        {{ formData.hostname }}
                </el-descriptions-item>
                <el-descriptions-item label="IP地址">
                        {{ formData.ip }}
                </el-descriptions-item>
                <el-descriptions-item label="管理IP">
                        {{ formData.manage_ip }}
                </el-descriptions-item>
                <el-descriptions-item label="网段ID">
                        {{ formData.network_id }}
                </el-descriptions-item>
                <el-descriptions-item label="管理网段ID">
                        {{ formData.manage_network_id }}
                </el-descriptions-item>
                <el-descriptions-item label="硬件配置模版">
                        {{ filterDict(formData.hardware_id,intOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="PXE模版">
                        {{ filterDict(formData.pxe_id,intOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="Kickstart模版">
                        {{ filterDict(formData.kickstart_id,intOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="机柜ID">
                        {{ filterDict(formData.cabinet_id,intOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="机柜U位">
                        {{ formData.cabinet_u }}
                </el-descriptions-item>
                <el-descriptions-item label="设备状态">
                        {{ formData.status }}
                </el-descriptions-item>
                <el-descriptions-item label="开发负责人">
                        {{ formData.manager_dev }}
                </el-descriptions-item>
                <el-descriptions-item label="运维负责人">
                        {{ formData.manager_ops }}
                </el-descriptions-item>
                <el-descriptions-item label="设备标签ID">
                        {{ filterDict(formData.label_id,intOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="设备描述">
                        {{ formData.device_describe }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createBareMetal,
  deleteBareMetal,
  deleteBareMetalByIds,
  updateBareMetal,
  findBareMetal,
  getBareMetalList
} from '@/api/sysDeviceBareMetal'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'BareMetal'
})

// 自动化生成的字典（可能为空）以及字段
const intOptions = ref([])
const formData = ref({
        sn: '',
        asset_id: '',
        hostname: '',
        ip: '',
        manage_ip: '',
        network_id: 0,
        manage_network_id: 0,
        hardware_id: undefined,
        pxe_id: undefined,
        kickstart_id: undefined,
        cabinet_id: undefined,
        cabinet_u: '',
        status: '',
        manager_dev: '',
        manager_ops: '',
        label_id: undefined,
        device_describe: '',
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
               asset_id : [{
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
               manage_ip : [{
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
               manage_network_id : [{
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
               cabinet_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               cabinet_u : [{
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
               status : [{
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
               manager_dev : [{
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
               manager_ops : [{
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
               label_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               device_describe : [{
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
  const table = await getBareMetalList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    intOptions.value = await getDictFunc('int')
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
            deleteBareMetalFunc(row)
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
      const res = await deleteBareMetalByIds({ ids })
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
const updateBareMetalFunc = async(row) => {
    const res = await findBareMetal({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reDeviceBareMetal
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBareMetalFunc = async (row) => {
    const res = await deleteBareMetal({ ID: row.ID })
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
  const res = await findBareMetal({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reDeviceBareMetal
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          sn: '',
          asset_id: '',
          hostname: '',
          ip: '',
          manage_ip: '',
          network_id: 0,
          manage_network_id: 0,
          hardware_id: undefined,
          pxe_id: undefined,
          kickstart_id: undefined,
          cabinet_id: undefined,
          cabinet_u: '',
          status: '',
          manager_dev: '',
          manager_ops: '',
          label_id: undefined,
          device_describe: '',
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
        asset_id: '',
        hostname: '',
        ip: '',
        manage_ip: '',
        network_id: 0,
        manage_network_id: 0,
        hardware_id: undefined,
        pxe_id: undefined,
        kickstart_id: undefined,
        cabinet_id: undefined,
        cabinet_u: '',
        status: '',
        manager_dev: '',
        manager_ops: '',
        label_id: undefined,
        device_describe: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createBareMetal(formData.value)
                  break
                case 'update':
                  res = await updateBareMetal(formData.value)
                  break
                default:
                  res = await createBareMetal(formData.value)
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
