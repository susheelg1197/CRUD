<template>
  <div class="container-fluid mt-5">
    <h1>Blood Bank Donor's List</h1>
    <div class="row">
      <div class="col-md-8">
        <table class="table table-striped">
          <thead>
            <tr>
              <th scope="col">Sr.No.</th>
              <th scope="col">Full Name</th>
              <th scope="col">Blood Group</th>
              <th scope="col">Email Id</th>
              <th scope="col">Action</th>
              <th scope="col">Verify</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(blood,index) in bloodDetails" :key="index">
              <th>{{index+1}}</th>
              <td><img v-if="blood.imageUrl" :src="blood.imageUrl" alt="Avatar" class="avatar mr-2">
              <img v-else src="https://gentle-bayou-82093.herokuapp.com/file/profilePictures/avatar.jpg" alt="Avatar" class="avatar mr-2">{{blood.firstName+' '}}{{blood.lastName}}</td>
              <td>{{blood.bloodGroup}}</td>
              <td>{{blood.email}}</td>
              <td>
          <b-icon-pencil @click="editInfo(blood)"></b-icon-pencil>
           <b-icon-trash class="ml-4" @click="deleteInfo(blood)"></b-icon-trash>
          </td>
          <td><button v-if="!blood.isVerified" type="button" class="btn btn-link" @click="verify(blood)">click here</button>
          <span v-else class="badge badge-primary ml-1">Verified</span>
           </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="col-md-4">
        <div class="card">
          <div class="card-header">
            What's new?
          </div>
          <div class="card-body">
            <h5 class="card-title">Donor's Registration</h5>
            <p class="card-text">
              Willing to donate blood? Please complete this Registration.
            </p>
            <a href="/#/donor-form" class="btn btn-primary">Click Here</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import axios from 'axios';
export default {
  data() {
    return {
      bloodDetails:[]
    }
  },
  methods: {
    editInfo(blood){
        this.$router.push({
        name:'EditDonorForm',
        params:{
          formNo:blood.formNo
        }
      })
    },
    getDetails(){
      axios.get('https://gentle-bayou-82093.herokuapp.com/o/blood-bank/show-details')
        .then((response) => {
                this.bloodDetails = response.data;
                });
    },
    deleteInfo(blood){
      var data={
        formNo:blood.formNo
      }
           axios.post('https://gentle-bayou-82093.herokuapp.com/o/blood-bank/delete-user-details',data)
            .then((response) => {
                console.log(response.data)
                });
              this.getDetails()
    },
    verify(blood){
      this.$router.push({
        name:'ImageUpload',
        params:{
          formNo:blood.formNo
        }
      })

    }
  },
  created() {
    axios.get('https://gentle-bayou-82093.herokuapp.com/o/blood-bank/show-details')
        .then((response) => {
                this.bloodDetails = response.data;
                });
}
}
</script>
<style>
.avatar {
  vertical-align: middle;
  width: 25px;
  height: 25px;
  border-radius: 50%;
}
.hero {
  height: 90vh;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
}
.hero .lead {
  font-weight: 200;
  font-size: 1.5rem;
}
</style>
