<template>
  <div class="flex p-10 space-x-4">
  <fwb-input
    v-model="setcode"
    placeholder="Enter set code (ex: THB)"
    label="Set Code"
    name="SetCode"
  />
  <fwb-input
    v-model="quantity"
    placeholder="Enter quantity"
    label="Quantity"
    name="Quantity"
  />
  <fwb-input
    v-model="cardnumber"
    placeholder="Ex: 123"
    label="Card Number"
    name="CardNumber"
  />
  <fwb-button type="submit" color="dark">Go!</fwb-button>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { FwbInput } from 'flowbite-vue'
import { FwbButton } from 'flowbite-vue'

const setcode = ref('')
const quantity = ref('0')
const cardnumber = ref('')

function getQuantity(): number {
  return Number(quantity.value)
}
</script>