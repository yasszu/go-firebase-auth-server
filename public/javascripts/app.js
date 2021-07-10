// Print signed in user info
firebase.auth().onAuthStateChanged(function (user) {
    if (user) {
        console.log('User is signed in.')
        user.getIdToken(/* forceRefresh */ true).then(function (idToken) {
            console.log(`idToken: ${idToken}`);
            axios.get('/v1/me', {
                headers: {
                    Authorization: `Bearer ${idToken}`,
                }
            })
                .then(function (response) {
                    // handle success
                    console.log(response.data);
                })
                .catch(function (error) {
                    // handle error
                    console.log(error);
                })

        }).catch(function (error) {
            console.log(error)
        });
    } else {
        console.log('No user is signed in.')
    }
});

// FirebaseUI config.
var uiConfig = {
    signInSuccessUrl: '/',
    callbacks: {
        signInSuccessWithAuthResult: function (authResult, redirectUrl) {
            if (authResult.user) {
                authResult.user.getIdToken(/* forceRefresh */ true).then(function (idToken) {
                    // Send token to your backend via HTTPS
                    var form = new FormData();
                    form.append('id_token', idToken);
                    axios.post('/authenticate', form)
                        .then(function (response) {
                            window.location.href = '/';
                        })
                        .catch(function (error) {
                            window.location.href = '/';
                        });
                }).catch(function (error) {
                    window.location.href = '/';
                });
                return false;
            }
            return true;
        },
        signInFailure: function (error) {
            // Some unrecoverable error occurred during sign-in.
            // Return a promise when error handling is completed and FirebaseUI
            // will reset, clearing any UI. This commonly occurs for error code
            // 'firebaseui/anonymous-upgrade-merge-conflict' when merge conflict
            // occurs. Check below for more details on this.
            return handleUIError(error);
        },
    },
    signInFlow: 'popup',
    signInOptions: [
        // Leave the lines as is for the providers you want to offer your users.
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
        firebase.auth.FacebookAuthProvider.PROVIDER_ID,
        firebase.auth.TwitterAuthProvider.PROVIDER_ID,
        firebase.auth.GithubAuthProvider.PROVIDER_ID,
        firebase.auth.EmailAuthProvider.PROVIDER_ID,
        firebase.auth.PhoneAuthProvider.PROVIDER_ID,
        firebaseui.auth.AnonymousAuthProvider.PROVIDER_ID
    ],
    // tosUrl and privacyPolicyUrl accept either url string or a callback
    // function.
    // Terms of service url/callback.
    tosUrl: '<your-tos-url>',
    // Privacy policy url/callback.
    privacyPolicyUrl: function () {
        window.location.assign('<your-privacy-policy-url>');
    }
};

// Initialize the FirebaseUI Widget using Firebase.
var ui = new firebaseui.auth.AuthUI(firebase.auth());

// The start method will wait until the DOM is loaded.
ui.start('#firebaseui-auth-container', uiConfig);
