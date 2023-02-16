function Serie1(){
    let nombre_saisons = 4;
    let nombre_episode_par_saison = 12;
    let duree_episode = 30; 
    let temps_de_pub = 2;

    let resultat1 = ((duree_episode + temps_de_pub) * nombre_episode_par_saison ) * nombre_saisons

    document.getElementById("Serie1()").innerHTML = resultat1;

}

function Serie2(){
    let nombre_saisons = 10;
    let nombre_episode_par_saison = 24;
    let duree_episode = 60; 
    let temps_de_pub = 2;

    let resultat2 = ((duree_episode + temps_de_pub) * nombre_episode_par_saison ) * nombre_saisons
    
    document.getElementById("Serie2()").innerHTML = resultat2;
}

function Serie3(){
    let nombre_saisons = 6;
    let nombre_episode_par_saison = 20;
    let duree_episode = 45; 
    let temps_de_pub = 2;

    let resultat3 = ((duree_episode + temps_de_pub) * nombre_episode_par_saison ) * nombre_saisons

    document.getElementById("Serie3()").innerHTML = resultat3;
}
