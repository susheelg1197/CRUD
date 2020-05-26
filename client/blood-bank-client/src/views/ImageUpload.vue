<template>
  <div class="container">
    <div class="row">
      <div class="col-lg-7 mx-auto">
        <div class="bg-white rounded-lg shadow-sm p-5">
          <div class="tab-content">
            <h1 class="text-info"><strong>Verification</strong></h1>
            <p><em>Donors Document Information</em></p>
            <div id="nav-tab-card" class="tab-pane fade show active">
              <form role="form" @submit.prevent="confirmUpload()">
                <div class="card mb-3">
                  <div class="card-header bg-transparent border-danger">
                    Profile Picture
                  </div>
                  <div class="card-body">
                    <div class="text-center mb-3">
                      <img
                        :src="currentUrl"
                        alt="Image Not Found"
                        class="img-thumbnail"
                        style="height:250px;width:250px"
                      />
                    </div>
                    <label>
                      <input
                        type="file"
                        id="file"
                        ref="file"
                        v-on:change="handleFileUpload()"
                      />
                    </label>
                    <button type="button" v-on:click="submitFile()">Submit</button>
                  </div>
                </div>
                <button
                  type="submit"
                  class="subscribe btn btn-primary btn-block rounded-pill shadow-sm"
                >
                  Confirm
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  /*
      Defines the data used by the component
    */
  data() {
    return {
      formNo: this.$route.params.formNo,
      file: "",
      currentUrl: "https://gentle-bayou-82093.herokuapp.com/file/profilePictures/avatar.jpg",
    };
  },

  methods: {
    confirmUpload(){
      var data={
        formNo:this.formNo,
        imageUrl:"https://gentle-bayou-82093.herokuapp.com/file/profilePictures/"+this.formNo+".png"
      }
      axios
      .post("https://gentle-bayou-82093.herokuapp.com/o/blood-bank/update-document-upload",data)
        .then(function() {
          console.log("SUCCESS!!");
        })
        .catch(function() {
          console.log("FAILURE!!");
        });
    },
    submitFile() {
      let formData = new FormData();
      console.log(this.file);
      formData.append("file", this.file);
      axios
        .post("https://gentle-bayou-82093.herokuapp.com/one/" + this.formNo, formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        })
        .then(function() {
          console.log("SUCCESS!!");
        })
        .catch(function() {
          console.log("FAILURE!!");
        });
    },

    /*
        Handles a change on the file upload
      */
    handleFileUpload() {
      this.file = this.$refs.file.files[0];
      this.currentUrl = URL.createObjectURL(this.file);
    },
  },
};
</script>
