<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-lg-7 mx-auto">
        <div class="bg-white rounded-lg shadow-sm p-5">
          <div class="tab-content">
            <h1 class="text-danger"><strong>Give Blood Save Life</strong></h1>
            <p><em>Donors Registration Form</em></p>
            <div id="nav-tab-card" class="tab-pane fade show active">
              <form role="form" @submit.prevent="confirm()">
                <div class="card mb-3">
                  <div class="card-header bg-transparent border-danger">
                    Full Name 
                  </div>
                  <div class="card-body">
                    <div class="row">
                      <div class="col-md-6">
                        <div class="form-group">
                          <label for="username">First Name</label>
                          <input
                            v-model="firstName"
                            type="text"
                            name="username"
                            placeholder="Susheel"
                            required
                            class="form-control"
                          />
                        </div>
                      </div>
                      <div class="col-md-6">
                        <div class="form-group">
                          <label for="username">Last Name</label>
                          <input
                            v-model="lastName"
                            type="text"
                            name="username"
                            placeholder="Gounder"
                            required
                            class="form-control"
                          />
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="card mb-3">
                  <div class="card-header bg-transparent border-danger">
                    Address
                  </div>
                  <div class="card-body">
                    <div class="form-group">
                      <label for="username">Street Address</label>
                      <input
                        v-model="address.streetAddress"
                        type="text"
                        name="username"
                        placeholder="Flat No 786"
                        required
                        class="form-control"
                      />
                    </div>
                    <div class="row">
                      <div class="col-md-6">
                        <div class="form-group">
                          <label for="username">City</label>
                          <input
                            v-model="address.city"
                            type="text"
                            name="username"
                            placeholder="Pune"
                            required
                            class="form-control"
                          />
                        </div>
                      </div>
                      <div class="col-md-6">
                        <div class="form-group">
                          <label for="username">State</label>
                          <input
                            v-model="address.state"
                            type="text"
                            name="username"
                            placeholder="Maharashtra"
                            required
                            class="form-control"
                          />
                        </div>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-6">
                        <div class="form-group">
                          <label for="username">PinCode</label>
                          <input
                            v-model="address.pincode"
                            type="text"
                            name="username"
                            placeholder="411018"
                            required
                            class="form-control"
                          />
                        </div>
                      </div>
                      <div class="col-md-6">
                        <div class="form-group">
                          <label for="username">Country</label>
                          <input
                            v-model="address.country"
                            type="text"
                            name="username"
                            placeholder="India"
                            required
                            class="form-control"
                          />
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="card mb-3">
                  <div class="card-header bg-transparent border-danger">
                    Blood Group
                  </div>
                  <div class="card-body">
                    <div class="form-group">
                      <b-form-select
                        class="mb-3"
                        v-model="selectedBloodGroup"
                        placeholder="Please"
                        :options="bloodGroupArray"
                      >
                      </b-form-select>
                    </div>
                  </div>
                </div>
                <div class="card mb-3">
                  <div class="card-header bg-transparent border-danger">
                    Contact Information
                  </div>
                  <div class="card-body">
                    <div class="form-group">
                      <label for="username">Phone Number</label>
                      <input
                        v-model="phoneNumber"
                        type="text"
                        name="username"
                        placeholder="9875527371"
                        required
                        class="form-control"
                      />
                    </div>
                    <div class="form-group">
                      <label for="username">Email Address </label>
                      <input
                        v-model="emailId"
                        type="email"
                        name="username"
                        placeholder="abcd@yourmail.com"
                        required
                        class="form-control"
                      />
                    </div>
                  </div>
                </div>
                 <div class="card mb-3">
                  <div class="card-header bg-transparent border-danger">
                    Feedback
                  </div>
                  <div class="card-body">
                    <div class="form-group">
                      <label for="username">Your Valuable Feedback</label>
                       <textarea v-model="feedback" class="form-control" id="exampleFormControlTextarea1" rows="3"></textarea>
                    </div>
                   
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
          <!-- End -->
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
      selectedBloodGroup: "Please Select Your Blood Group",
      bloodGroupArray: [
        "Please Select Your Blood Group",
        "A+",
        "A-",
        "B+",
        "B-",
        "O+",
        "O-",
        "AB+",
        "AB-",
      ],
      address: {
        city: "",
        streetAddress: "",
        state: "",
        country: "",
        pincode: "",
      },
      firstName: "",
      lastName: "",
      phoneNumber: "",
      emailId: "",
      feedback: "",
      formCount:0,
      info:null,
      bloodDetails:[]
    };
  },
  methods: {
      confirm(){
          this.formCount=this.formCount.toString()
          var data={
              formNo:this.$route.params.formNo,
              firstName:this.firstName,
              lastName:this.lastName,
              address:this.address,
              phoneNumber:this.phoneNumber,
              email:this.emailId,
              feedback:this.feedback,
              bloodGroup:this.selectedBloodGroup
          }
          console.log(data)
        axios
      .post('https://gentle-bayou-82093.herokuapp.com/o/blood-bank/update-user-details',data)
      .then(response => {
          console.log(response.data)
          this.$router.push('/')
      })
      }
  },
   created() {
       var bloodObj ={}
    axios.get('https://gentle-bayou-82093.herokuapp.com/o/blood-bank/show-details')
        .then((response) => {
                this.bloodDetails = response.data;
                bloodObj=response.data.filter(a=>a.formNo===this.$route.params.formNo)[0]
                 this.firstName=bloodObj.firstName
                 this.lastName=bloodObj.lastName
                 this.address=bloodObj.address
                 this.selectedBloodGroup=bloodObj.bloodGroup
                 this.phoneNumber=bloodObj.phoneNumber
                 this.emailId=bloodObj.email
                 this.feedback=bloodObj.feedback
                });
               
}
};
</script>
<style scoped>
body {
  background: #f5f5f5;
}

.rounded-lg {
  border-radius: 1rem;
}

.nav-pills .nav-link {
  color: #555;
}

.nav-pills .nav-link.active {
  color: #fff;
}
</style>
