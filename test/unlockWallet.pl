#!  /usr/bin/perl -w

################################################################################
#   requires the installation of cpan, and after needs to run the following:
#   $ echo "install Expect" | cpan
################################################################################

use strict;
use Expect;

my  $PLDCTL = './bin/pldctl';
my  $PLDCTL_ARGS = '';
my  $WALLET_PASSWORD = 'w4ll3tP@sswd';
my  $WALLET_PUBLIC_PASSWORD = '';

#   check cli arguments
my  $VERBOSE = '';

if( $#ARGV + 1 == 1 ) {
    $VERBOSE = $ARGV[ 0 ];
}

#   open a pipe to interact with pldctl
my  $pldctl = Expect->new;
my  $inputLine;

$pldctl->raw_pty( 1 );
$pldctl->spawn( qq/$PLDCTL $PLDCTL_ARGS unlock/ )
    or die qq/Failed to open pipe to $PLDCTL: $!/;

if( $VERBOSE eq '--verbose' ) {
    $pldctl->exp_internal( 1 );
}

#   wait pldctl ask for wallet private password, and the send it
$pldctl->expect( 3, 'Input wallet private password: ' );
$pldctl->send( qq/$WALLET_PASSWORD\n/ );

#   wait pldctl ask for wallet public password, and the send it
$pldctl->expect( 3, 'Input wallet public password: ' );
$pldctl->send( qq/$WALLET_PUBLIC_PASSWORD\n/ );

#   wait pldctl shows a successful message
$pldctl->expect( 3, 'lnd successfully unlocked!' );
