<template>
  <div>
    <div class="toolbar">
      <el-date-picker v-model="dateRange" type="daterange" range-separator="至"
        start-placeholder="开始日期" end-placeholder="结束日期" style="width:300px" />
      <el-select v-model="filterStatus" placeholder="状态" clearable style="width:110px">
        <el-option label="成功" value="成功" />
        <el-option label="失败" value="失败" />
      </el-select>
      <el-button :icon="Download" @click="exportOrders">导出</el-button>
    </div>

    <el-table :data="orders" stripe border style="width:100%;margin-top:16px" size="small">
      <el-table-column prop="id" label="订单号" width="140" />
      <el-table-column prop="machine" label="设备" min-width="160" />
      <el-table-column prop="prize" label="奖品" min-width="120" />
      <el-table-column prop="user" label="用户" width="100" />
      <el-table-column prop="amount" label="金额" width="80" align="center">
        <template #default="{ row }">
          <span style="color:#6366f1;font-weight:600">¥{{ row.amount }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="80" align="center">
        <template #default="{ row }">
          <el-tag :type="row.status === '成功' ? 'success' : 'danger'" size="small">{{ row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="time" label="时间" width="160" />
    </el-table>

    <div style="margin-top:16px;display:flex;justify-content:flex-end">
      <el-pagination background layout="total,prev,pager,next" :total="orders.length" :page-size="10" />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Download } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const dateRange = ref([])
const filterStatus = ref('')

const orders = [
  { id: 'ORD-20240001', machine: '商场A-01号机', prize: '皮卡丘公仔', user: 'u_****23', amount: '5.00', status: '成功', time: '2024-03-11 14:32:10' },
  { id: 'ORD-20240002', machine: '商场B-03号机', prize: '蜡笔小新', user: 'u_****87', amount: '5.00', status: '成功', time: '2024-03-11 14:28:55' },
  { id: 'ORD-20240003', machine: '商场A-02号机', prize: '限定盲盒', user: 'u_****41', amount: '10.00', status: '失败', time: '2024-03-11 14:25:03' },
  { id: 'ORD-20240004', machine: '地铁站-01号机', prize: '熊本熊公仔', user: 'u_****66', amount: '5.00', status: '成功', time: '2024-03-11 14:20:18' },
  { id: 'ORD-20240005', machine: '商场C-01号机', prize: '奥特曼玩具', user: 'u_****09', amount: '5.00', status: '成功', time: '2024-03-11 14:15:44' },
  { id: 'ORD-20240006', machine: '商场B-01号机', prize: '史努比钥匙扣', user: 'u_****12', amount: '3.00', status: '失败', time: '2024-03-11 14:10:22' },
  { id: 'ORD-20240007', machine: '商场A-01号机', prize: '哆啦A梦抱枕', user: 'u_****55', amount: '5.00', status: '成功', time: '2024-03-11 13:58:01' },
]

function exportOrders() {
  ElMessage.success('正在导出订单数据…')
}
</script>

<style scoped>
.toolbar { display: flex; gap: 12px; align-items: center; flex-wrap: wrap; }
</style>
