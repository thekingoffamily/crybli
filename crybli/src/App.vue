<template>
  <div>
    <img src="../public/img/icons/crybli.jpg" class="logo" alt="logo" />
    <main>
      <p>Crybli update's very soon...</p>
      <br>
      <br>{{ this.btcPrice }}
      <br>
      <p>Contract: 0x5f69330c89B157CcCf042B138DC8cf4aa6F6ceA2</p>
    </main>
    <footer>
      <div class="white-line"></div>
        <a href="https://github.com/thekingoffamily/crybli/tree/master/crybli" target="_blank">Project on GitHub</a>
        <!-- <p style="color:aliceblue;">Contract: 0x5f69330c89B157CcCf042B138DC8cf4aa6F6ceA2</p> -->
        <a href="https://bscscan.com/token/0x5f69330c89B157CcCf042B138DC8cf4aa6F6ceA2" target="_blank">Contract on BscScan</a>
        <a href="https://pancakeswap.finance/swap?outputCurrency=0x5f69330c89B157CcCf042B138DC8cf4aa6F6ceA2&chainId=56&inputCurrency=BNB" target="_blank">Buy in PancakeSwap</a>
    </footer>
  </div>
</template>

<script>

export default({
  data(){
    return{
      btcPrice: '',
    }
  },
  methods:{
    getBtcPrice(){
      this.btcPrice = fetch('http://localhost:8080/getBtcPrice');
    },
  },

  mounted(){
    setInterval(()=>{this.getBtcPrice();}, 10000);
  }

})

</script>

<style>

  body {
    background-color: #000;
  }
  /* Стили для белой полосы в футере */
  .white-line {
    /* height: 3px;
    background-color: #fff; */
  }

  /* Стили для ссылки на GitHub репозиторий */
  footer a {
    display: block;
    text-align: center;
    margin-top: 10px;
    font-size: 18px;
  }

  /* Стили для логотипа в хедере */
  .logo {
    display: block;
    margin: 20px auto;
    width: 150px;
    height: 150px;
  }

  main {
    color: aliceblue;
    text-align: center;
    height: 50vh;
    font-size: 22px;
  }
</style>
