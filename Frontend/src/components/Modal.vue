<template >
    <div class="modal-backdrop" @click="close">
      <div class="modal">
        <header class="modal-header">
          <span class="modal__title">{{title}}</span>
          <img class="btn-close" @click="close"  src="@/assets/close.svg" alt="">
        </header>
        <section class="modal-body">
          <slot name="body" 
          ></slot>
        </section>
        <footer class="modal-footer">
          <slot name="footer"
          ></slot>
        </footer>
      </div>
    </div>
  </template>
  
  <script>
    export default {
      name: 'ModalPage',
      props: {
        title: {
          type: String,
          required: false
        }
      },
      methods: {
        close() {
          this.$emit('close')
        }
      }
    };
  </script>
  
  <style lang="scss">
    .modal-backdrop {
      position: fixed;
      z-index: 999;
      top: 0;
      bottom: 0;
      left: 0;
      right: 0;
      background-color: rgba(0, 0, 0, 0.3);
      display: flex;
      justify-content: center;
      align-items: center;
    }
    .modal {
      position: relative;
      padding: 30px 30px 40px 30px;
      width: 630px;
      z-index: 1000;
      background: #FFFFFF;
      display: flex;
      flex-direction: column;
      border-radius: 8px;
      &__title {
        font-weight: 600;
        font-size: 24px;
        display: block;
        color: var(--black-title);
        margin-bottom: 30px;
      }
      &-footer {
        margin-top: 30px;
        display: flex;
        align-items: center;
      }
    }
    .btn-close {
      position: absolute;
      height: 20px;
      top: 17px;
      right: 17px;
      border: none;
      cursor: pointer;
      font-weight: bold;
      color: #2196F3;
      background: transparent;
    }
  </style>