$(document).ready(function() {
	window.onload = function() {
  	$(document.getElementById("mm1")).fadeOut(100)
    $(document.getElementById("mmm1")).fadeOut(100)
    $(document.getElementById("mm2")).fadeOut(100)
    $(document.getElementById("mmm2")).fadeOut(100)
    $(document.getElementById("mm3")).fadeOut(100)
    $(document.getElementById("mmm3")).fadeOut(100)
    $(document.getElementById("mm4")).fadeOut(100)
    $(document.getElementById("mmm4")).fadeOut(100)
    $(document.getElementById("mm5")).fadeOut(100)
    $(document.getElementById("mmm5")).fadeOut(100)
        
  }
  
    // Mobile Menu Toggle
    $('.mobile-menu-toggle').on('click', function(e) {
        e.preventDefault()
        e.stopPropagation()
        $('.mobile-menu').show()
    })
    $('body').on('click', function() {
        $('.mobile-menu').hide()
    })

    // Reset initial header height
    $('header').height(window.innerHeight - 60)

    // Home Page Scroll Buttons
    $('.scroll-down, .get-started').on('click', function(e) {
        e.preventDefault()
        $('html, body').animate({ scrollTop: $('section:first-of-type').offset().top }, 300)
    })

    // Sidebar Accordian
    $('.sidebar-category').on('click', function() {
        $(this).toggleClass('active')
    })
    
    
    
    
    // Modal Display1
    $(document.getElementById("m1")).on('click', function() {
        $(document.getElementById("mm1")).fadeIn(200)
        $(document.getElementById("mmm1")).fadeIn(200)
    })
    $('.modal-mask, .modal-save').on('click', function() {
        $(document.getElementById("mm1")).fadeOut(100)
        $(document.getElementById("mmm1")).fadeOut(100)
    })
    
    // Modal Display2
    $(document.getElementById("m2")).on('click', function() {
        $(document.getElementById("mm2")).fadeIn(200)
        $(document.getElementById("mmm2")).fadeIn(200)
    })
    $('.modal-mask, .modal-save').on('click', function() {
        $(document.getElementById("mm2")).fadeOut(100)
        $(document.getElementById("mmm2")).fadeOut(100)
    })
    
    
   /*  $(document.getElementById("m3")).on('click', function() {
        $('.modal-mask2, .modal').fadeIn(200)
    })
    $('.modal-mask2, .modal-save').on('click', function() {
        $('.modal-mask2, .modal').fadeOut(100)
    })
    */
    
    // Modal Display3
     $(document.getElementById("m3")).on('click', function() {
        $(document.getElementById("mm3")).fadeIn(200)
        $(document.getElementById("mmm3")).fadeIn(200)
    })
    $('.modal-mask, .modal-save').on('click', function() {
        $(document.getElementById("mm3")).fadeOut(100)
        $(document.getElementById("mmm3")).fadeOut(100)
    })
    
    // Modal Display4
     $(document.getElementById("m4")).on('click', function() {
        $(document.getElementById("mm4")).fadeIn(200)
        $(document.getElementById("mmm4")).fadeIn(200)
    })
    $('.modal-mask, .modal-save').on('click', function() {
        $(document.getElementById("mm4")).fadeOut(100)
        $(document.getElementById("mmm4")).fadeOut(100)
    })
    
    // Modal Display5
     $(document.getElementById("m5")).on('click', function() {
        $(document.getElementById("mm5")).fadeIn(200)
        $(document.getElementById("mmm5")).fadeIn(200)
    })
    $('.modal-mask, .modal-save').on('click', function() {
        $(document.getElementById("mm5")).fadeOut(100)
        $(document.getElementById("mmm5")).fadeOut(100)
    })



    // Tabs
    $('.tab').on('click', function(e) {
        e.preventDefault()
        $(this).closest('.tabs').find('.tab').removeClass('active')
        $(this).addClass('active')
    })
})
