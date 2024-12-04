
<template>
    <DarkMode />
    <div class="containerAndButtons-profile">
      <div class="container-profile">
        <label class="circle-label-profile"/>
        <div class="info-profile">
          <div class="infoUser-profile">
            <p class="p-profile" id="Username-profile" ></p>
            <div class="LastAndFirstName-profile">
              <p class="p-profile" id="First-Name"></p>
              <p class="p-profile" id="Last-Name"></p>
            </div>
            <p class="p-profile" id="Birthday"></p>
            <p class="p-profile" id="Email"></p>
            <p class="p-profile" id="About"></p>
          </div>
          <div class="FollowAndFollower-profile">
            <div class="buttons-profile">
              <router-link v-if="isOwner" id="signUP" to="/settings">
                <button class="edit-profile" type="submit">Edit</button>
              </router-link>
              
              <button v-else-if="isPrivate && !isFollow && !isWaiting" @click="sendrequestprivate('')" class="follow-profile-button-private">send a request</button>
              <button v-else-if="isPrivate && !isFollow && isWaiting" @click="sendunrequestprivate('')" class="follow-profile-button-waiting">waiting a response</button>
              <button v-else-if="!isFollow" @click="followRequest('')" class="follow-profile-button">Follow</button>
              <button v-else="isFollow" @click="unFollowRequest('')" class="unfollow-profile-button">UnFollow</button>
              <button class="share-profile">Share</button>
            </div>
            <div class="stats-profile align-horizontal">
              <div @click="toggleVisibleFollowersList('followers')" class=" align-horizontal follower-following-profile gap-followers ">
                <p id="Followers"></p>
                <p>followers</p>
              </div>
              <div id="followersList" v-show="followersListIsVisibile" class="align-vertical">
                  <div id="modal-followers-lists" class="lists-modals-profilepage modal-followers">
                    <div class="search-bar-container">
                      <input v-model="searchQuery" type="text" placeholder="search a user..."class="search-bar"/>
                    </div>
                    <div class="of-followers "v-for="follower in filteredFollowers">
                        <img @click="redirectionToProfile(follower.Id)" :src="follower.Pictures">
                        <p @click="redirectionToProfile(follower.Id)"> {{ follower.Username }}</p>
                      <button @click="unFollowRequest(follower.Id)"v-if="follower.HeIsFollowing && follower.Id != userState.id" class="unfollow-profile-button">unfollow</button>
                      <button @click="sendrequestprivate(follower.Id)"v-else-if="!follower.HeIsWaiting && follower.Privacy =='Private' && follower.Id != userState.id" class="follow-profile-button-private">send a request</button>
                      <button @click="sendunrequestprivate(follower.Id)"v-else-if="follower.HeIsWaiting && follower.Privacy && follower.Id != userState.id"class="follow-profile-button-waiting" >waiting</button>
                      <button @click="followRequest(follower.Id)"v-else-if="follower.Privacy=='Public' && !follower.HeIsFollowing  && follower.Id != userState.id" class="follow-profile-button">follow</button>
                      <button @click="removeTheUser(follower.Id)" v-if="typeOfFollow=='followers' && isOwner">remove</button>
                    </div>
            
                  </div>
                </div>
              <div @click="toggleVisibleFollowersList('following')" class="align-horizontal follower-following-profile gap-followers">
                <p id="Following"></p>
                <p>Following</p>
              
              </div>
              <div v-if="isOwner && isPrivate"class="follower-input">
                <div @click="toggleVisibleFollowersWaitingList"id="number-follower-inbox"></div>
                <div id="modal-followers-waiting" v-show="followersWaitingListIsVisible">
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  <script>
    import { Settings } from "../services/userServices";
   import { inject  } from "vue";
   import router from '../router';
   import { followRequestJs, unFollowRequestJs, requestFollowJs, unRequestFollowJs, fetchWaitingFollowers, fetchFollowersFromTheProfile, removeTheUserInTheList} from "../services/profileServices";
   export default {
    props: ["idUser"],
    data() {
        return {
        userState: inject("userState"),
        follow: false,
        privacy: false,
        waiting: false,
        followersWaitingListIsVisible: false,
        followersListIsVisibile: false,
        allfollowers: null,
        searchQuery: '',
        typeOfFollow: null
        };
    },
    computed: {
      filteredFollowers() {
          if (this.allfollowers) {
            return this.allfollowers.filter(follower =>
              follower.Username.toLowerCase().includes(this.searchQuery.toLowerCase())
            );
          }
          return [];
        },
        isOwner() {
        return this.userState.id === this.idUser;
        },
        isFollow(){
          return this.follow == true
        },
        isPrivate(){
          return this.privacy == "Private"
        },
        isWaiting(){
          return this.waiting == true
        }
    },
    mounted() {
         Settings.handleProfile(this).then(result=>{
          this.follow = result.isfollow
          this.privacy = result.privacy
          this.waiting = result.iswaiting
        })
    }, 
    methods: {

      async toggleVisibleFollowersList(whatsWeLookingfor){
        this.typeOfFollow = whatsWeLookingfor
        this.allfollowers = await fetchFollowersFromTheProfile(this.idUser, this.userState.id, whatsWeLookingfor)
        this.followersListIsVisibile = !this.followersListIsVisibile
      },

      removeTheUser(iduser){
        removeTheUserInTheList(this, iduser);
      },

      toggleVisibleFollowersWaitingList(){
        this.followersWaitingListIsVisible = !this.followersWaitingListIsVisible
        fetchWaitingFollowers(this.userState.id)
      },
       followRequest(userToFollow){
         followRequestJs(this.userState.id, userToFollow!="" ? userToFollow: this.idUser, this, userToFollow=="" ? false: true).then(result=>{ this.follow = result })
        },
        unFollowRequest(userToUnFollow){
          unFollowRequestJs(this.userState.id, userToUnFollow!="" ? userToUnFollow: this.idUser, this, userToUnFollow=="" ? false: true, this.typeOfFollow).then(result=>{ this.follow = !result })
        },
        sendrequestprivate(userToFollow){
          requestFollowJs(this.userState.id, userToFollow!="" ? userToFollow: this.idUser, this,userToFollow=="" ? false: true).then(result=>{ this.waiting = result })
        },
        sendunrequestprivate(userToUnFollow){
          unRequestFollowJs(this.userState.id, userToUnFollow!="" ? userToUnFollow: this.idUser, this,userToUnFollow=="" ? false: true).then(result=>{ this.waiting = !result })
        },
        redirectionToProfile(idToRedirect){
          router.push(`/${idToRedirect}`)
        }    
    }
    };
</script>